package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	solve_B(solve_A())
}

var whole_range = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127}
var col_range = []int{0, 1, 2, 3, 4, 5, 6, 7}

func solve_A() []int {
	ids := make([]int, 0)
	boarding_passes := readFile("input.txt")
	my_range := whole_range
	my_col_range := col_range
	seat_id := 0
	for _, pass := range boarding_passes {

		rows := pass[:7]
		columns := pass[7:]

		r := strings.SplitAfter(rows, "")
		c := strings.SplitAfter(columns, "")

		for _, row := range r {
			if row == "F" {
				my_range = my_range[:len(my_range)/2]
			} else if row == "B" {
				my_range = my_range[len(my_range)/2:]
			}
		}
		row_value := my_range[0]

		for _, col := range c {
			if col == "L" {
				my_col_range = my_col_range[:len(my_col_range)/2]
			} else if col == "R" {
				my_col_range = my_col_range[len(my_col_range)/2:]
			}
		}
		col_value := my_col_range[0]

		id := row_value*8 + col_value
		//fmt.Println("ID:", id)
		ids = append(ids, id)

		if id > seat_id {
			seat_id = id
		}

		my_range = whole_range
		my_col_range = col_range

	}

	fmt.Println("Seat ID:", seat_id)
	fmt.Println("LEN:", len(ids))
	sort.Ints(ids)
	return ids
}

func solve_B(ids []int) {
	N := len(ids)
	b := make([]int, ids[N-1])

	for i := 0; i < N-1; i++ {
		b[ids[i]] = 1
	}

	for i := ids[0]; i <= ids[N-1]-1; i++ {
		if b[i] == 0 {
			fmt.Print(i, " ")
		}
	}

}

func readFile(input string) []string {
	content, err := ioutil.ReadFile(input)
	check(err)

	text := string(content)
	slice := strings.Split(text, "\n")

	if len(slice) > 0 {
		slice = slice[:len(slice)-1]
	}

	return slice
}

func check(e error) { // check if error == nil
	if e != nil {
		fmt.Println(e)
	}
}
