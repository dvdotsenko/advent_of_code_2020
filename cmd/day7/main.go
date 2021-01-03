// https://golang.org/ref/spec
package main

import (
	"aoc2020/pkg/day7parser"
	"aoc2020/pkg/fio"
	"fmt"
)

const FN string = "cmd/day7/data.txt"


func main() {
	lines := fio.FileAsLines(FN)
	for line := range lines {
		if len(line) > 0 {
			fmt.Printf("> %s\n", line)
			fmt.Printf("< %v\n", day7parser.Parse(line))
		}
	}
}
