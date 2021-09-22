package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	name  string
	sign  string
	value int
	first bool
}

func main() {
	Solve_B("input.txt")
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
		Instructions = append(Instructions, Instruction{instruction, sign, value, true})
		//fmt.Printf("%s -> %s:%d\n", instruction, sign, value)
		i++
		x++
	}

	j := 0
	//z := 0

	b := true
	for true {

		instr := Instructions[j].name
		sgn := Instructions[j].sign
		val := Instructions[j].value

		//fmt.Printf("%s %s%d  | %d\n", instr, sgn, val, z+1)

		if Instructions[j].first {
			if sgn == "+" {
				if instr == "nop" {
					Instructions[j].first = false
					j++
				} else if instr == "acc" {
					Instructions[j].first = false
					counter += val
					j++
				} else if instr == "jmp" {
					Instructions[j].first = false
					j += val
				}
			} else if sgn == "-" {
				if instr == "nop" {
					Instructions[j].first = false
					j++
				} else if instr == "acc" {
					counter -= val
					Instructions[j].first = false
					j++
				} else if instr == "jmp" {
					Instructions[j].first = false
					j -= val
				}
			}
		} else if !Instructions[j].first {
			b = false
			break
		}

		_ = b

		//z++
	}
	fmt.Println("Counter:", counter)
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
