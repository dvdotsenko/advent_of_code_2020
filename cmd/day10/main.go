// https://golang.org/ref/spec
package main

import (
	"aoc2020/pkg/fio"
	"aoc2020/pkg/set"
	"errors"
	"fmt"
	"strconv"
)

type Any = interface{}

const EmptyString = ""

var OutOfBoundsError = errors.New("Out of bounds.")

func GetNumbers(fn string) []int64 {
	lines := fio.FileAsLines(fn)
	nn := make([]int64, 0, 200)
	for line := range lines {
		if line != EmptyString {
			n, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				panic(err)
			}
			// OyO
			nn = append(nn, n)
		}
	}
	return nn
}

func FindBadSum(nn []int64, window int) (sum int64, i int, err error) {

	l := len(nn)
	var w set.IntSet
	for i := 0 + window; i < l; i += 1 {
		sum := nn[i]
		w = set.NewIntSet(nn[i-window : i]...)
		fmt.Printf("%v: %v, %v\n", i, w, sum)

		found_reciprocal := false
		for v := range w {
			reciprocal_v := sum - v
			if w.Contains(reciprocal_v) {
				found_reciprocal = true
				break
			}
		}

		if !found_reciprocal {
			fmt.Printf("Did not find components for %v in %v (i: %v)\n", sum, w, i)
			return sum, i, nil
		}
	}
	return 0, 0, errors.New("Anomaly not found")
}

func CheckRange(sum int64, nn []int64) (int, error) {
	var cum int64 = 0
	for i, v := range nn {
		cum += v
		if cum == sum {
			return i, nil
		} else if cum > sum {
			return 0, OutOfBoundsError
		}
	}
	return 0, OutOfBoundsError
}

func FindMiddle(i1 int, i2 int) int {
	// inclusive indexes in array.
	// (3, 4) (3, 5) diff 2, middle 4
	// (3, 5) (3, 6) diff 3, middle 4
	// (3, 6) (3, 7) diff 4, middle 5
	// (3, 7) (3, 8) diff 5, middle 5
	return ((i2+1)-i1)/2 + i1
}

func FindRange(sum int64, nn []int64) (int, int, error) {
	l := len(nn)
	start := 0
	for start < l {
		end, err := CheckRange(sum, nn[start:])
		if err != nil {
			start += 1
			continue
		}
		return start, start + end, nil
	}
	return 0, 0, OutOfBoundsError
}

func main() {

	nn := GetNumbers("cmd/day10/data_example.txt")

	var max int64
	for _, v := range nn {
		if v > max {
			max = v
		}
	}

	nno := make([]int64, max+1, max+1)
	for _, v := range nn {
		nno[v] = v
	}

	var v int64

	l := len(nno)
	s := 0
	s1 := 0
	s2 := 0
	s3 := 1 // last jump guaranteed
	for i := 1; i < l; i += 1 {
		v = nno[i]
		if v == 0 {
			s += 1
		} else {
			if s == 0 {
				s1 += 1
			} else if s == 1 {
				s2 += 1
			} else if s == 2 {
				s3 += 1
			} else {
				fmt.Printf("Count %v found and cannot be handled", s)
				panic("No!")
			}
			s = 0
		}
	}

	fmt.Printf("s1: %v, s2: %v, s3: %v, factorized: %v\n", s1, s2, s3, s1*s3)

	mm := make(map[int]int)
	cnt := 0
	nno[0] = 100000
	for i := 0; i < l; i += 1 {
		v = nno[i]
		if v == 0 {
			if cnt > 0 {
				mm[cnt] += 1
			}
			cnt = 0
		} else {
			cnt += 1
		}
	}
	if cnt > 0 {
		mm[1] += 1
	}

	permutation_multiplier := map[int]int{
		//1: 0,
		//2: 0,
		3: 2,
		4: 4,
		5: 8 - 1, // 8 permutations, but 1, where all digits are missing, is not allowed because gap will be > 3
	}

	var aa int = 1
	for k, v := range mm {
		if k > 2 {
			aa *= permutation_multiplier[k] * v
		}
	}

	fmt.Printf("For %v, permutations are %v\n", mm, aa)

}
