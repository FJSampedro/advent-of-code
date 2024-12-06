package main

import (
	"fmt"
	"lib"
	"regexp"
	"strconv"
)

func partOne() {
	// Reading file
	lines, err := lib.ReadTxtLines()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Regex commands
	r, err := regexp.Compile(`mul\(\d*,\d*\)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	var commands []string
	for _, line := range lines {
		commands = append(commands, r.FindAllString(line, -1)...)
	}

	//Iter commands
	r2, err := regexp.Compile(`\d+`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	var sum int
	for _, command := range commands {

		multipliers := r2.FindAllString(command, -1)
		m1, err := strconv.Atoi(multipliers[0])
		if err != nil {
			fmt.Println("Error converting first multiplier:", err)
			return
		}
		m2, err := strconv.Atoi(multipliers[1])
		if err != nil {
			fmt.Println("Error converting second multiplier::", err)
			return
		}
		sum += m1 * m2
	}
	fmt.Println(sum)
}

func partTwo() {
	// Reading file
	lines, err := lib.ReadTxtLines()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Regex commands
	r, err := regexp.Compile(`mul\(\d*,\d*\)|do\(\)|don't\(\)`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	var commands []string
	for _, line := range lines {
		commands = append(commands, r.FindAllString(line, -1)...)
	}

	//Iter commands
	r2, err := regexp.Compile(`\d+`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	var sum int
	active := true

	for _, command := range commands {
		if command == "do()" {
			active = true
			continue
		}
		if command == "don't()" {
			active = false
			continue
		}
		if active {
			multipliers := r2.FindAllString(command, -1)
			m1, err := strconv.Atoi(multipliers[0])
			if err != nil {
				fmt.Println("Error converting first multiplier:", err)
				return
			}
			m2, err := strconv.Atoi(multipliers[1])
			if err != nil {
				fmt.Println("Error converting second multiplier::", err)
				return
			}
			sum += m1 * m2
		}
	}
	fmt.Println(sum)
}
func main() {
	fmt.Println("################## Part One ##################")
	partOne()
	fmt.Println("################## Part Two ##################")
	partTwo()
}
