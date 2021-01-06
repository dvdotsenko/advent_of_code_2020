// https://golang.org/ref/spec
package main

import (
	"aoc2020/pkg/day7parser"
	"aoc2020/pkg/fio"
	"fmt"
)

const FN string = "cmd/day7/data.txt"

type Any = interface{}

const EmptyString = ""

//func main() {
//	lines := fio.FileAsLines(FN)
//	mm := map[string][]day7parser.Component{}
//	reverse_mm := map[string]map[string]bool{}
//	for line := range lines {
//		if len(line) > 0 {
//			fmt.Printf("> %s\n", line)
//			expr := day7parser.Parse(line).(day7parser.Expr)
//			fmt.Printf("< %v\n", expr)
//			mm[expr.Label] = expr.Components
//			for _, c := range expr.Components {
//				_, exists := reverse_mm[c.Label]
//				if !exists {
//					reverse_mm[c.Label] = map[string]bool{}
//				}
//				reverse_mm[c.Label][expr.Label] = true
//			}
//		}
//	}
//
//	for k, vv := range reverse_mm {
//		for v, _ := range vv {
//			fmt.Printf("+ %v: %v\n", k, v)
//		}
//	}
//
//	cc := map[string]bool{}
//	s := "shiny gold"
//	for p, _ := range reverse_mm[s] {
//		cc[p] = true
//	}
//
//	pass := 1
//	for {
//		cnt := 0
//		for c, _ := range cc {
//			fmt.Printf("%v: > '%s': %v\n", pass, c, cc)
//			vv, exists := reverse_mm[c]
//			if exists {
//				for p, _ := range vv {
//					if _, already_there := cc[p]; already_there {
//						fmt.Printf("%v: '%s' is already there\n", pass, p)
//					} else {
//						cnt += 1
//						cc[p] = true
//						fmt.Printf("%v: '%s' adding\n", pass, p)
//					}
//				}
//			}
//		}
//		if cnt == 0 {
//			break
//		}
//		pass += 1
//	}
//	fmt.Printf("%v, %v", len(cc), cc)
//}

func SumValues(s string, mm map[string]map[string]int64) int64 {
	var cc int64 = 1
	m := mm[s]
	for k, v := range m {
		cc += SumValues(k, mm) * v
	}
	return cc
}

func main() {
	lines := fio.FileAsLines(FN)
	mm := map[string]map[string]int64{}
	for line := range lines {
		if len(line) > 0 {
			fmt.Printf("> %s\n", line)
			expr := day7parser.Parse(line).(day7parser.Expr)
			fmt.Printf("< %v\n", expr)

			mm[expr.Label] = map[string]int64{}
			for _, c := range expr.Components {
				mm[expr.Label][c.Label] = c.Count
			}
		}
	}

	////fmt.Print(mm)
	//
	//js, err := json.Marshal(mm)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fio.ToFile(js, "./cmd/day7/data.json")

	//var i int64 = 0
	//for v := range NodesParentAdjustedValues("shiny gold", mm) {
	//	fmt.Printf("=== Value Consumed: %v\n", v)
	//	i += v
	//}
	//fmt.Print(i)

	//mm := map[string]map[string]int64{}
	//js := fio.FromFile("./cmd/day7/data.json")
	//json.Unmarshal(js, &mm)

	i := SumValues("shiny gold", mm)
	print(i)
}
