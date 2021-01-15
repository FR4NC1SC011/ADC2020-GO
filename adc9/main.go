package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input := readFile("example.txt")
	solve_A(input)

}

func solve_A(input []string) {

	for i := 5; i < len(input); i++ {
		sumOfInput(input, 5, i)

	}

}

func sumOfInput(input []string, preamble int, start int) {
	sum_array := input[start : start+preamble]
	fmt.Println(sum_array)
}

func validNumber(validNumbers []int, number int) bool {

	return false
}

func readFile(input string) []string {
	content, err := ioutil.ReadFile(input)
	Check(err)

	text := string(content)
	s := strings.Split(text, "\n")

	return s
}

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
