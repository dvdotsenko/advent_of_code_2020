# https://adventofcode.com/2020/day/1

# focus on solving the scan problem in as few iterations as possible (without going nuts)

from math import factorial

test_data = [
    1721,
    979,
    366,
    299,
    675,
    1456,
]


data = list(map(int, """1511
1344
1925
1970
1864
1951
1557
1984
1743
1526
1972
1945
1969
1760
2008
1592
736
1963
1994
2009
1777
1856
1899
1926
1850
687
2005
1094
1949
1326
2002
1805
1493
1341
1828
1778
1767
1364
1973
1768
1929
1377
2000
1726
1913
2001
1574
1859
1793
1957
1959
1388
1593
1392
724
1962
1999
252
1982
1662
1892
1610
1343
1831
1862
1991
1394
1946
1935
1986
1911
1358
1322
1956
1988
1758
1490
1998
1744
1844
1294
1764
1543
1560
1562
1747
1870
1292
1989
1752
1471
1980
1897
1544
1914
1923
1944
1375
1987
1993
1742
1975
1479
1977
1934
1939
1950
1992
1983
1474
1643
2010
1814
1942
322
1425
1646
1878
1410
1927
1761
1948
1779
1753
1847
274
1659
1773
1960
1772
1674
1809
1568
1978
1952
1947
1976
1953
1961
1937
1932
1781
2007
1941
1393
1573
1745
169
89
1408
1974
1810
1979
1967
890
1958
1930
1954
1759
720
1936
1576
1407
2004
1964
1462
1875
1943
1938
2006
1739
1378
1922
1924
2003
1792
1985
1729
1966
1355
1940
1928
1357
1955
1896
1115
1836
1971
1329
1807
1997
1359
1801
1933
1965
1981
1711
1905
1625
1968""".splitlines()))


N = 2020


get_midpoint = lambda a, b: a + int((b - a) / 2)


def check(a, b, N, nn):
    v = nn[a] + nn[b]
    if v == N:
        return '='
    if v < N:
        return '<'
    if v > N:
        return '>'


get_combinations = lambda l, k: factorial(l) / (
        factorial(k) * factorial(l - k)
)


def find_2_numbers_eq_N(nn: list[int], N):
    print('comparing to ', N)
    nn = list(sorted(nn))
    # find position on list that is smaller than N

    x = -1
    l = len(nn)
    cnt = 0

    combinations = get_combinations(l, 2)

    while x < l:
        x += 1
        a = x
        b = l-1
        tried = set()
        tried.add(x)
        while True:
            y = get_midpoint(a, b)
            if y in tried:
                break
            tried.add(y)

            cnt += 1
            if cnt > combinations:
                raise Exception('Epic fail. It took you more iterations than there are combinations')

            print(f'{cnt}: x={x} ({nn[x]}) + y={y} ({nn[y]}) = ({nn[x]+nn[y]}) ( a={a}, b={b} )')
            if y == a:
                # range exhausted. Only 2 elements and first one selected
                # try end of range (because it's never selected because round-down by int())
                y = b
                c = check(x, y, N, nn)
                if c == '=':
                    return nn[x], nn[y], cnt
                # print(f'{cnt}: {c} x={x} ({nn[x]}), y={y} ({nn[y]}), ({nn[x]+nn[y]}) a={a}, b={b}')
                break # while. goes to next x
            else: # y > x
                c = check(x, y, N, nn)
                if c == '=':
                    return nn[x], nn[y], cnt
                elif c == '>':
                    # print(f'{cnt}: {c} x={x} ({nn[x]}), y={y} ({nn[y]}), ({nn[x]+nn[y]}) a={a}, b={b}')
                    b = y
                elif c == '<':
                    # print(f'{cnt}: {c} x={x} ({nn[x]}), y={y} ({nn[y]}), ({nn[x]+nn[y]}) a={a}, b={b}')
                    a = y

    return None, None, cnt


def doit_for_2(nn):
    x, y, cnt = find_2_numbers_eq_N(nn, 2020)
    if x is None or y is None:
        raise Exception("No Value")
    c = get_combinations(len(nn), 2)
    print(f'Values: {x}+{y}={x+y}, {x}*{y}={x*y}  with {cnt} tries out of {c} possible')


def doit_for_3(nn):
    nn = list(sorted(nn))
    l = len(nn)
    i = -1
    c = get_combinations(l, 3)
    total_cnt = 0
    while i < l:
        i += 1
        x = nn[i]

        y, z, cnt = find_2_numbers_eq_N(nn[:i]+nn[i+1:], 2020-x)

        total_cnt += cnt
        if y is None or z is None:
            continue # while
        print(f'Values: {x}i+{y}+{z}={x+y+z}, {x}*{y}*{z}={x*y*z} with {total_cnt} tries out of {c} possible')
        break


if __name__ == '__main__':
    doit_for_3(data)
    # 169i+736+1115=2020, 169*736*1115=138688160 with 1204 tries out of 1313400 possible combinations
    print('=====================================')
    doit_for_2(data)
    # 252+1768=2020, 252*1768=445536  with 23 tries out of 19900 possible combinations
