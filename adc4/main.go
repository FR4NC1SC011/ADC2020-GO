package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//	solve_A()
	solve_B()
}

func solve_A() {
	var input []string
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	var required int
	var valid_passport int
	input = readFile("input.txt")

	for _, passport := range input {
		for _, field := range fields {
			if strings.Contains(passport, field) {
				required++
			}
		}

		if required == len(fields) {
			valid_passport++
		} else if required == len(fields)-1 && strings.Contains(passport, "cid") != true {
			valid_passport++
		}
		required = 0
	}

	fmt.Println("Valid Passports:", valid_passport)
}

func solve_B() {
	valids := 0
	input := readFile("input.txt")

	for _, passport := range input {
		passportFields := make(map[string]string)
		re := regexp.MustCompile("\\n")
		passport = re.ReplaceAllString(passport, " ")

		for _, field := range strings.Fields(passport) {
			passportFields[field[:3]] = field[4:]
		}
		//fmt.Println(passportFields)

		_, hasCid := passportFields["cid"]
		if len(passportFields) == 8 || len(passportFields) == 7 && !hasCid {
			valid := verifyRange(passportFields["byr"], 1920, 2002) && verifyRange(passportFields["iyr"], 2010, 2020) && verifyRange(passportFields["eyr"], 2020, 2030)
			valid = valid && verifyHeight(passportFields["hgt"]) && verifyHair(passportFields["hcl"]) && verifyEye(passportFields["ecl"]) && verifyPid(passportFields["pid"])
			if valid {
				valids++
				fmt.Println(passportFields)
			}
		}

	}
	fmt.Println("Valid:", valids)
}

func readFile(input string) []string {
	content, err := ioutil.ReadFile(input)
	check(err)

	text := string(content)
	s := strings.Split(text, "\n\n")

	return s
}

func check(e error) { // check if error == nil
	if e != nil {
		fmt.Println(e)
	}
}

func verifyRange(field string, min, max int) bool {
	value, _ := strconv.Atoi(field)
	return value >= min && value <= max
}

func verifyHeight(field string) bool {
	pattern := "((1[5-8][0-9])|(19[0-3]))cm|((59)|(6[0-9])|(7[0-6]))in"
	valid, _ := regexp.MatchString(pattern, field)
	return valid
}

func verifyHair(field string) bool {
	if len(field) == 7 {
		pattern := "#(([0-9])|([a-f]))"
		valid, _ := regexp.MatchString(pattern, field)
		return valid
	}
	return false
}

func verifyEye(field string) bool {
	pattern := "(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)"
	valid, _ := regexp.MatchString(pattern, field)
	return valid
}

func verifyPid(field string) bool {
	pattern := "^\\d{9}$"
	valid, _ := regexp.MatchString(pattern, field)
	return valid

}
