package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	second_part()
}

func readFile(input string) []string {
	content, err := ioutil.ReadFile(input)
	check(err)

	text := string(content)
	s := strings.Split(text, "\n")
	//fmt.Println(s)

	return s

}

func pass_rule(password string) (int, int, string, string) {
	policy := password[:strings.IndexByte(password, ':')]

	pass := password[strings.IndexByte(password, ':')+2:]
	min := policy[:strings.IndexByte(policy, '-')]
	max := policy[strings.IndexByte(policy, '-')+1 : strings.IndexByte(policy, ' ')]
	letter := policy[strings.IndexByte(policy, ' ')+1:]

	minimum, err := strconv.Atoi(min)
	check(err)
	maximum, err := strconv.Atoi(max)
	check(err)

	return minimum, maximum, letter, pass

}

func first_part() {
	var valid_passwords int
	passwords := readFile("input.txt")
	fmt.Println("size:", cap(passwords))

	for i := 1; i < len(passwords)-1; i++ {
		min, max, letter, passwd := pass_rule(passwords[i])
		count := strings.Count(passwd, letter)
		if count >= min && count <= max {
			valid_passwords++
		}

	}
	fmt.Println("Valid Passwords:", valid_passwords)
}

func second_part() {
	var valid_passwords int
	passwords := readFile("input.txt")

	for i := 1; i < len(passwords)-1; i++ {
		first, second, letter, passwd := pass_rule(passwords[i])
		first_index := passwd[first-1 : first]
		second_index := passwd[second-1 : second]

		if first_index == letter && second_index != letter {
			valid_passwords++
			continue
		} else if first_index != letter && second_index == letter {
			valid_passwords++
			continue
		}

	}
	fmt.Println("Valid Passwords:", valid_passwords)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
