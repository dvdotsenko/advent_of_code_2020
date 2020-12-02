import re
from collections import defaultdict

line_pattern = re.compile('(?P<min_cnt>\d+)-(?P<max_cnt>\d+)\s+(?P<ch>\w+):\s+(?P<text>\w+)')


test_data = '''1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc'''.splitlines()

with open('day2.data') as fp:
    data = fp.readlines()


def is_valid_password_sled_rental(min_cnt, max_cnt, ch, text):
    # print(min_cnt, max_cnt, ch, text)
    mi = int(min_cnt)
    ma = int(max_cnt)
    cnt = defaultdict(int)
    for l in text:
        cnt[l] += 1
    return cnt[ch] >= mi and cnt[ch] <= ma


def is_valid_password_tobaggan(min_cnt, max_cnt, ch, text):
    print(min_cnt, max_cnt, ch, text)
    p1 = int(min_cnt) - 1
    p2 = int(max_cnt) - 1
    s = {text[p1], text[p2]}
    if ch in s and len(s) == 2:
        return True
    else:
        return False


def is_valid_line(line, rule_fn):
    m = re.match(line_pattern, line)
    if not m:
        raise ValueError(f'Line "{line}" does not have data matching required pattern')
    return rule_fn(**m.groupdict())


print(
    'good passwords - task 1:',
    sum([
        1 if is_valid_line(line, is_valid_password_sled_rental) else 0
        for line in data
    ])
)

print(
    'good passwords - task 2:',
    sum([
        1 if is_valid_line(line, is_valid_password_tobaggan) else 0
        for line in data
    ])
)
