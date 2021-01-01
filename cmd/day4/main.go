// https://golang.org/ref/spec
package main

import (
	"aoc2020/pkg/fio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const FN string = "cmd/day4/data.txt"


type RecordData map[string]string


func doNothing (s string) bool {
	return true
}


func validateNumber(s string, min int64, max int64) bool {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}
	if v < min || v > max {
		return false
	} else {
		return true
	}
}


var _hairColors = map[string]interface{}{
	"amb":nil,"blu":nil,"brn":nil,"gry":nil,"grn":nil,"hzl":nil,"oth":nil,
}

func validateEyeColor(s string) bool {
	_, there := _hairColors[s]
	return there
}


var hairColor = regexp.MustCompile(`^#[a-f0-9]{6}$`)
var validPassportID = regexp.MustCompile(`^[0-9]{9}$`)


func validatePassport(s string) bool {
	return validPassportID.MatchString(s)
}


func validateHair(s string) bool {
	return hairColor.MatchString(s)
}


var gg = regexp.MustCompile(`^(\d+)(cm|in)$`)


func validateHeight(s string) bool {
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	if !gg.MatchString(s) {
		return false
	}
	vv := gg.FindStringSubmatch(s)
	fmt.Print(vv, "\n")
	if vv[2] == "cm" {
		return validateNumber(vv[1], 150, 193)
	}
	if vv[2] == "in" {
		return validateNumber(vv[1], 59, 76)
	}
	return false
}


var FIELD_VALIDATOR_MAP = map[string]func(string)bool{
	"byr": func (s string) bool {return validateNumber(s, 1920, 2002)}, // (Birth Year)
	//"cid", // (Country ID)
	"ecl": validateEyeColor, // (Eye Color)
	"eyr": func (s string) bool {return validateNumber(s, 2020, 2030)}, // (Expiration Year)
	"hcl": validateHair, // (Hair Color)
	"hgt": validateHeight, // (Height)
	"iyr": func (s string) bool {return validateNumber(s, 2010, 2020)}, // (Issue Year)
	"pid": validatePassport, // (Passport ID)
}


func parseRecord(s string) RecordData {
	record := map[string]string{}
	var x []string
	for _, part := range strings.Split(s, " ") {
		if part != "" {
			x = strings.SplitN(part, ":", 2)
			record[x[0]] = x[1]
		}
	}
	return record
}


func getRecords (fileLines chan string) chan RecordData {
	var records = make(chan RecordData, 1)
	go func () {
		var parts []string
		for line := range fileLines {
			if line == "" {
				records <- parseRecord(strings.Join(parts, " "))
				parts = parts[0:0]
			} else {
				parts = append(parts, line)
			}
		}
		close(records)
	}()
	return records
}


func recordIsValid(r RecordData) bool {
	for k, validate := range FIELD_VALIDATOR_MAP {
		v, there := r[k]
		if !there {
			return false
		}
		if !validate(v) {
			return false
		}
	}
	return true
}


func main() {
	lines := fio.FileAsLines(FN)
	var cnt int
	for record := range getRecords(lines) {
		fmt.Print(">  ", record, "\n")
		if recordIsValid(record) {
			cnt += 1
		}
	}
	print(cnt)
}
