package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var replacer = strings.NewReplacer(" ", "")

func main() {
	solve_B()

}

func solve_A() {
	input, err := os.Open("example.txt")
	check(err)
	defer input.Close()

	sc := bufio.NewScanner(input)

	var answer string
	var count int

	for {
		end := !sc.Scan()

		if sc.Text() == "" || end {
			str := answer
			str = replacer.Replace(str)
			s := strings.Split(str, "")
			unique := unique_count(s)
			count += len(unique)
			answer = ""
		} else {
			answer += sc.Text() + " "
		}

		if end {
			break
		}
	}
	fmt.Println("Answer", count)
}

func solve_B() {
	input, err := os.Open("example.txt")
	check(err)
	defer input.Close()

	sc := bufio.NewScanner(input)

	var answer string
	var count int

	for {
		end := !sc.Scan()

		if sc.Text() == "" || end {
			if len(answer) > 4 {
				str := answer
				str = replacer.Replace(str)
				s := strings.Split(str, "")
				dup_map := duplicate_count(s)
				for _, v := range dup_map {
					//fmt.Printf("Item: %s, N: %d\n", k, v)
					if v >= 2 {
						count += 1
					}
				}
			} else {
				count += len(answer) - 1
			}
			answer = ""
		} else {
			answer += sc.Text() + " "
		}

		if end {
			break
		}
	}
	fmt.Println("Count:", count)

}

func unique_count(input []string) []string {
	keys := make(map[string]bool)
	output := []string{}
	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			output = append(output, entry)
		}
	}
	return output
}

func duplicate_count(input []string) map[string]int {
	duplicate_frequency := make(map[string]int)

	for _, item := range input {
		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1
		} else {
			duplicate_frequency[item] = 1
		}
	}
	return duplicate_frequency
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
