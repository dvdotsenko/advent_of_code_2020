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

	nn := GetNumbers("cmd/day9/data.txt")

	sum, i, err := FindBadSum(nn, 25)
	if err != nil {
		panic(err)
	}

	nn_short := nn[:i]

	i1, i2, err := FindRange(sum, nn_short)
	if err != nil {
		panic(err)
	}

	vv := nn[i1 : i2+1]
	var s int64
	var min int64 = sum
	var max int64
	for _, v := range vv {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		s += v
	}

	fmt.Printf("Indexes: %v:%v, values %v. sum: %v, min: %v, max: %v, control value: %v \n", i1, i2, vv, s, min, max, min+max)

}
