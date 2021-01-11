package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	name   string
	sign   string
	value  int
	repeat bool
}

func main() {
	solve_A("example.txt")
}

func solve_A(I string) {
	instructions := readFile(I)

	var counter int

	var instruction string
	var sign string
	var str_value string
	var value int

	var Instructions []Instruction

	x := 0
	i := 0
	for x < len(instructions)-1 {
		instruction = instructions[i]
		instruction = instruction[0:3]
		sign = instructions[i]
		sign = sign[4:5]
		str_value = instructions[i]
		str_value = str_value[5:]
		value, _ = strconv.Atoi(str_value)
		Instructions = append(Instructions, Instruction{instruction, sign, value, false})
		//fmt.Printf("%s -> %s:%d\n", instruction, sign, value)
		i++
		x++
	}

	j := 0
	z := 0
	for z < len(instructions)-1 {

		instr := Instructions[j].name
		sgn := Instructions[j].sign
		val := Instructions[j].value
		rep := Instructions[j].repeat

		if !rep {
			fmt.Println("Entramos")
			if sgn == "+" {
				if instr == "jmp" {
					j += val
				} else if instr == "acc" {
					counter += val
				} else if instr == "nop" {
					j++
				}
			} else if sgn == "-" {
				if instr == "jmp" {
					j -= val
				} else if instr == "acc" {
					counter -= val
				} else if instr == "nop" {
					j++
				}
			}
			fmt.Printf("%s -> %s:%d\n", Instructions[j].name, Instructions[j].sign, Instructions[j].value)
			Instructions[j].repeat = true
		} else if rep {
			fmt.Println(counter)
			return
		}
		z++
	}
}

func readFile(input string) []string {
	content, err := ioutil.ReadFile(input)
	check(err)

	text := string(content)
	s := strings.Split(text, "\n")

	return s
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
