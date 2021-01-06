// https://golang.org/ref/spec
package main

import (
	"aoc2020/pkg/day8parser"
	"aoc2020/pkg/fio"
	"errors"
	"fmt"
	"sort"
)

const FN string = "cmd/day8/data.txt"

type Any = interface{}

const EmptyString = ""

const (
	ACC  = "acc"
	JMP  = "jmp"
	NOP  = "nop"
	STOP = "stop"
)

var LoopError = errors.New("Infinite Loop")
var OutOfBoundsError = errors.New("Out of Bounds")

func CheckAlgo(ops []day8parser.Expr) (int64, error, map[int64]bool) {
	var op day8parser.Expr
	var i int64 = 0
	var l int64 = int64(len(ops))
	var v int64 = 0
	seen := map[int64]bool{}

	for i > -1 && i < l {
		if seen[i] {
			print("LOOP DETECTED. STOPPING\n")
			return 0, LoopError, seen
		}
		seen[i] = true
		op = ops[i]
		fmt.Printf("sum: %v i: %v op: %v\n", v, i, op)
		if op.Op == ACC {
			v += op.Value
			i += 1
		}
		if op.Op == NOP {
			i += 1
		}
		if op.Op == JMP {
			i += op.Value
		}
	}

	if i == l {
		return v, nil, seen
	} else {
		return 0, OutOfBoundsError, seen
	}
}

func FlipOp(i int, ops []day8parser.Expr) []day8parser.Expr {
	new_ops := make([]day8parser.Expr, len(ops))
	l := copy(new_ops, ops)
	if l != len(ops) {
		panic("Copy right things, damn it.")
	}
	op := new_ops[i]

	if op.Op == JMP {
		new_ops[i] = day8parser.Expr{
			NOP, op.Value,
		}
	} else if op.Op == NOP {
		new_ops[i] = day8parser.Expr{
			JMP, op.Value,
		}
	} else {
		panic("Must be one of JMP or NOP")
	}
	return new_ops
}

func main() {
	lines := fio.FileAsLines(FN)
	ops := make([]day8parser.Expr, 0, 200)
	for line := range lines {
		if len(line) > 0 {
			expr := day8parser.Parse(line).(day8parser.Expr)
			ops = append(ops, expr)
			//fmt.Printf("< %v\n", expr)
		}
	}
	fmt.Printf("< %v\n", ops)

	v, err, seen := CheckAlgo(ops)
	fmt.Printf("%v out of %v ops used\n", len(seen), len(ops))

	doChanges := false
	if errors.Is(err, LoopError) || errors.Is(err, OutOfBoundsError) {
		doChanges = true
	}

	if !doChanges {
		print(v)
		return
	}

	// must study and change
	print("Analyzing the issue...\n")
	ii := make([]int, 0, len(ops))
	for i := range seen {
		if op := ops[i]; op.Op == NOP || op.Op == JMP {
			ii = append(ii, int(i))
		}
	}
	sort.Ints(ii)
	fmt.Print(ii)

	for _, i := range ii {
		ops_new := FlipOp(i, ops)
		v, err, _ := CheckAlgo(ops_new)
		if errors.Is(err, LoopError) || errors.Is(err, OutOfBoundsError) {
			continue
		} else {
			fmt.Printf("Flipped Op %v and got sum %v\n", i, v)
			break
		}
	}
}
