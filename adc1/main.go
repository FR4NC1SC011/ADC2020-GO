package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	i := readFile("input.txt")
	fmt.Println(i)

	solve(i)
}

func solve(input []int) int {
	for _, x := range input {
		for _, y := range input[1:] {
			for _, z := range input[2:] {
				if x+y+z == 2020 {
					fmt.Println("Encontrado:", x, y, z)
					output := x * y * z
					fmt.Println("Respuesta:", output)
				}
			}

		}
	}
	return 0

}

func readFile(input string) []int {
	var n []int
	content, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	s := strings.Split(text, "\n")
	for _, v := range s {
		i, _ := strconv.Atoi(v)
		n = append(n, i)
	}
	return n
}
