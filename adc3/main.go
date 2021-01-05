package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		input, err := os.Open("input.txt")
		check(err)

		defer input.Close()

		sc := bufio.NewScanner(input)
		var rightSteps, trees int
		for sc.Scan() {
			if sc.Text()[rightSteps%len(sc.Text())] == '#' { //I use module operator to prevent overflow
				trees++
			}
			rightSteps += 1
			fmt.Println(rightSteps, " % ", len(sc.Text()))
			fmt.Println("Modulo:", rightSteps%len(sc.Text()))
		}
		fmt.Println(trees)
	*/

	last_part()

}

func check(e error) {
	if e != nil {
		fmt.Println("You need a file called input.txt in this directory")
	}
}

func last_part() {
	var lines []string
	input, err := os.Open("input.txt")
	check(err)

	defer input.Close()

	sc := bufio.NewScanner(input)
	var rightSteps, trees int
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	for i := 0; i < len(lines); i += 2 {
		line := lines[i]
		if line[rightSteps%len(line)] == '#' {
			trees++
			fmt.Println("tree", trees)
		}
		rightSteps++

	}

}

// 1,1 = 80
// 3,1 = 162
// 5,1 = 77
// 7,1 = 83
// 1,2 = 37
