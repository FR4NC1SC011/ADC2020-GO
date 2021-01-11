package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bagLabel struct {
	bag    string
	number int
}

func main() {
	solve_B("ex2.txt")
}

func solve_A(input string) {
	rules := make(map[string][]string)
	valids := make(map[string]bool)
	valids["shiny gold"] = true

	file, err := os.Open(input)
	check(err)
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		parsedInput := strings.ReplaceAll(sc.Text(), " bags", "")
		parsedInput = strings.ReplaceAll(parsedInput, " bag", "")
		parsedInput = strings.ReplaceAll(parsedInput, ".", "")
		rule := strings.Split(parsedInput, " contain ")

		for _, bag := range strings.Split(rule[1], ", ") {
			rules[rule[0]] = append(rules[rule[0]], bag[2:])
		}
	}

	for bag := range rules {
		if lookIn(bag, rules, valids) {
			valids[bag] = true
		}
	}

	fmt.Println(len(valids) - 1)
}

func solve_B(input string) {
	rules := make(map[string][]bagLabel)

	file, err := os.Open(input)
	check(err)
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		parsedInput := strings.ReplaceAll(sc.Text(), " bags", "")
		parsedInput = strings.ReplaceAll(parsedInput, " bag", "")
		parsedInput = strings.ReplaceAll(parsedInput, ".", "")
		rule := strings.Split(parsedInput, " contain")

		for _, bag := range strings.Split(rule[1], ", ") {
			numberOfBags, _ := strconv.Atoi(bag[1:2])
			rules[rule[0]] = append(rules[rule[0]], bagLabel{bag[2:], numberOfBags})
		}
	}
	fmt.Println(rules)
	fmt.Println(countIn("shiny gold", rules) - 1)

}

func lookIn(lookedBag string, rules map[string][]string, valids map[string]bool) bool {
	if valids[lookedBag] {
		return true
	}

	for _, bag := range rules[lookedBag] {
		if lookIn(bag, rules, valids) {
			valids[bag] = true
			return true
		}
	}
	return false
}

func countIn(lookedBag string, rules map[string][]bagLabel) int {
	countBags := 1
	for _, bag := range rules[lookedBag] {
		countBags += bag.number * countIn(bag.bag, rules)
	}
	return countBags
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
