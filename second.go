package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solve_B() {
	input, err := os.Open("input.txt")
	check(err)

	defer input.Close()

	sc := bufio.NewScanner(input)
	var (
		passaport string
		valids    int
	)

	for {
		end := !sc.Scan()

		if sc.Text() == "" || end {
			passaportFields := make(map[string]string)
			for _, field := range strings.Fields(passaport) {
				passaportFields[field[:3]] = field[4:]
			}
			_, hasCid := passaportFields["cid"]

			if len(passaportFields) == 8 || len(passaportFields) == 7 && !hasCid {
				valid := verifyRange(passaportFields["byr"], 1920, 2002) && verifyRange(passaportFields["iyr"], 2010, 2020) && verifyRange(passaportFields["eyr"], 2020, 2030)
				valid = valid && verifyHeight(passaportFields["hgt"]) && verifyHair(passaportFields["hcl"]) && verifyEye(passaportFields["ecl"]) && verifyPid(passaportFields["pid"])

				if valid {
					valids++
				}

			}

			passaport = ""
		} else {
			passaport = sc.Text() + " "
			break
		}

		if end {
			break
		}
	}

	fmt.Println("Valid Passports:", valids)

}

func verifyRange(field string, min, max int) bool {
	value, _ := strconv.Atoi(field)
	return value >= min && max >= value
}

func verifyHeight(field string) bool {
	pattern := "((1[5-8][0-9])|(19[0-3]))cm|((59)|(6[0-9])|(7[0-6]))in"
	valid, _ := regexp.MatchString(pattern, field)
	return valid
}

func verifyHair(field string) bool {
	pattern := "#(([0-9])|([a-f]))"
	valid, _ := regexp.MatchString(pattern, field)
	return valid
}

func verifyEye(field string) bool {
	pattern := "^(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)$"
	valid, _ := regexp.MatchString(pattern, field)
	return valid
}

func verifyPid(field string) bool {
	pattern := "^\\d{9}$"
	valid, _ := regexp.MatchString(pattern, field)
	return valid
}
