package main

import (
	"aoc2023/days"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type DayFunction func(data []string, part int) int

var DAYS []DayFunction = []DayFunction{
	days.One,
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		println("Expected: aoc2023 <day>")
	}

	day, err := strconv.Atoi(argsWithoutProg[0])
	check(err)

	dat, err := os.ReadFile(fmt.Sprintf("data/%d/data.txt", day))
	check(err)
	data := strings.Split(string(dat), "\n")

	dat, err = os.ReadFile(fmt.Sprintf("data/%d/example-p1.txt", day))
	check(err)
	data_example_p1 := strings.Split(string(dat), "\n")

	dat, err = os.ReadFile(fmt.Sprintf("data/%d/example-p2.txt", day))
	check(err)
	data_example_p2 := strings.Split(string(dat), "\n")

	day_function := DAYS[day-1]

	fmt.Printf("P1: %d\n", day_function(data, 1))
	fmt.Printf("P2: %d\n", day_function(data, 2))
	fmt.Printf("P1 (Example): %d\n", day_function(data_example_p1, 1))
	fmt.Printf("P2 (Example): %d\n", day_function(data_example_p2, 2))
}
