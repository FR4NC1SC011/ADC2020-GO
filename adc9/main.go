package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input.txt")
	solve_A(input)

}

func solve_A(input []string) {
	var int_input []int

	for _, s_number := range input {
		number, _ := strconv.Atoi(s_number)
		int_input = append(int_input, number)
	}

	start := 0
	for i := 0; i < len(input); i++ {

		sums := sumOfInput(int_input, 25, i)

		if !validNumber(sums, int_input[i+25]) {
			fmt.Println("Invalid Number ->", int_input[i+25])
			break
		}
		start++
	}

}

func sumOfInput(input []int, preamble int, start int) []int {
	var sum_array []int
	preamble_array := input[start : start+preamble]

	for i := 0; i < len(preamble_array); i++ {
		for j := i + 1; j < len(preamble_array); j++ {
			sum := preamble_array[i] + preamble_array[j]
			//fmt.Println(preamble_array[i], "+", preamble_array[j], "=", sum)
			sum_array = append(sum_array, sum)

		}
	}

	return sum_array
}

func validNumber(validNumbers []int, number int) bool {

	for _, element := range validNumbers {
		if number == element {
			//fmt.Println("Valid Number -->", element)
			return true
		}
	}

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
