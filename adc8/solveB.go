package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	name    string
	value   int
	visited bool
}

func Solve_B(input string) {
	inp, err := os.Open(input)
	Check(err)
	defer inp.Close()

	sc := bufio.NewScanner(inp)

	var loop []instruction

	for sc.Scan() {
		line := strings.Fields(sc.Text())
		number, _ := strconv.Atoi(line[1])
		loop = append(loop, instruction{line[0], number, false})
	}

	for pCounter := range loop {
		if loop[pCounter].name == "acc" {
			continue
		}

		copyLoop := make([]instruction, len(loop))
		copy(copyLoop, loop)
	}

}

func Check(e error) {
	if e != nil {
		log.Println("Error")
	}

}
