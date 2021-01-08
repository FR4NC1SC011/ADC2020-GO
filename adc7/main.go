package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

type Bag struct {
	Name     string
	Quantity int
}

func main() {
	solve_A("example.txt")
}

func solve_A(input string) {
	var bag string
	var quantity int

	bags := make(map[string]Bag)
	var b Bag

	file, err := os.Open("example.txt")
	check(err)
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		if sc.Text() == "contain" {
			fmt.Println(bag)
			bag = ""
		} else if strings.HasSuffix(sc.Text(), ",") || strings.HasSuffix(sc.Text(), ".") {
			fmt.Println("  ", bag, quantity)
			b = Bag{
				Name:     bag,
				Quantity: quantity,
			}
			bag = ""
		} else if valid.IsInt(sc.Text()) {
			quantity, err = strconv.Atoi(sc.Text())
			check(err)
		} else if sc.Text() == "no" || sc.Text() == "other" {
			quantity = 0
			fmt.Println("    No")
		} else {
			bag += sc.Text() + " "
		}
		bags[bag] = b
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
