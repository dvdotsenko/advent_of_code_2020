// https://golang.org/ref/spec
package main

import (
	//"aoc2020/pkg/fio"
	"aoc2020/pkg/fio"
	"fmt"
	//"regexp"
	//"strconv"
	//"strings"
)

const FN string = "cmd/day5/data.txt"

var rows = [128]int{}
var seats = [8]int{}

func init() {
	for i := range rows {
		rows[i] = i
	}
	for i := range seats {
		seats[i] = i
	}
}

func pickLower(s []int) []int {
	l := len(s)
	return s[:l/2]
}

func pickUpper(s []int) []int {
	l := len(s)
	return s[l/2:]
}

var FIELD_PICKER_MAP = map[rune]func([]int) []int{
	'F': pickLower,
	'L': pickLower,
	'B': pickUpper,
	'R': pickUpper,
}

type RecordData struct {
	rowSteps  string
	seatSteps string
}

func getRecords(fileLines chan string) chan RecordData {
	var records = make(chan RecordData, 1)
	go func() {
		for line := range fileLines {
			if len(line) == 10 {
				records <- RecordData{
					line[:7],
					line[7:],
				}
			}
		}
		close(records)
	}()
	return records
}

func convertRecordToSeatIndex(r RecordData) int {
	var s []int
	s = rows[:]
	for _, ch := range r.rowSteps {
		s = FIELD_PICKER_MAP[ch](s)
		//fmt.Print(string(ch), " ", s, "\n")
	}
	row_num := s[0]

	//fmt.Print("Row number ", row_num, "\n")

	s = seats[:]
	for _, ch := range r.seatSteps {
		s = FIELD_PICKER_MAP[ch](s)
		//fmt.Print(string(ch), " ", s, "\n")
	}
	seat_num := s[0]
	//fmt.Print("Seat number ", seat_num, "\n")

	return row_num*8 + seat_num
}

func main() {
	lines := fio.FileAsLines(FN)
	var maxSeatIndex int
	var seatUsedMap = map[int]bool{}
	for record := range getRecords(lines) {
		seatIndex := convertRecordToSeatIndex(record)
		if seatIndex > maxSeatIndex {
			maxSeatIndex = seatIndex
		}
		fmt.Print(">  ", record, seatIndex, " ", maxSeatIndex, "\n")
		seatUsedMap[seatIndex] = true
	}
	for i := 0; i < 1024; i += 1 {
		if _, there := seatUsedMap[i]; !there {
			fmt.Printf("Missing Seat: %v\n", i)
		}
	}

	//fmt.Print(rows)
	//fmt.Print(seats)
}
