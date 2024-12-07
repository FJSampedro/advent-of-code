package main

import (
	"fmt"
	"lib"
)

func getChar(lines []string, px int, py int) (ch string) {
	if py >= 0 && px >= 0 && py < len(lines) && px < len(lines[py]) {
		ch = string(lines[py][px])
	} else {
		ch = "J"
	}
	return
}

func partOne() {
	lines, err := lib.ReadTxtLines()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	sum := 0
	for j, line := range lines {
		for i := range line {
			ch := getChar(lines, i, j)
			if ch == "X" {
				if getChar(lines, i-1, j-1) == "M" &&
					getChar(lines, i-2, j-2) == "A" &&
					getChar(lines, i-3, j-3) == "S" {
					sum++
				}

				if getChar(lines, i, j-1) == "M" &&
					getChar(lines, i, j-2) == "A" &&
					getChar(lines, i, j-3) == "S" {
					sum++
				}

				if getChar(lines, i+1, j-1) == "M" &&
					getChar(lines, i+2, j-2) == "A" &&
					getChar(lines, i+3, j-3) == "S" {
					sum++
				}

				if getChar(lines, i+1, j) == "M" &&
					getChar(lines, i+2, j) == "A" &&
					getChar(lines, i+3, j) == "S" {
					sum++
				}

				if getChar(lines, i+1, j+1) == "M" &&
					getChar(lines, i+2, j+2) == "A" &&
					getChar(lines, i+3, j+3) == "S" {
					sum++
				}
				if getChar(lines, i, j+1) == "M" &&
					getChar(lines, i, j+2) == "A" &&
					getChar(lines, i, j+3) == "S" {
					sum++
				}
				if getChar(lines, i-1, j+1) == "M" &&
					getChar(lines, i-2, j+2) == "A" &&
					getChar(lines, i-3, j+3) == "S" {
					sum++
				}
				if getChar(lines, i-1, j) == "M" &&
					getChar(lines, i-2, j) == "A" &&
					getChar(lines, i-3, j) == "S" {
					sum++
				}

			}
		}
	}
	fmt.Println(sum)
}

func partTwo() {
	lines, err := lib.ReadTxtLines()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	sum := 0
	for j, line := range lines {
		for i := range line {
			ch := getChar(lines, i, j)
			if ch == "A" {
				//M.S
				//.A.
				//M.S
				if getChar(lines, i-1, j-1) == "M" &&
					getChar(lines, i+1, j+1) == "S" &&
					getChar(lines, i-1, j+1) == "M" &&
					getChar(lines, i+1, j-1) == "S" {
					sum++
				}

				//M.M
				//.A.
				//S.S
				if getChar(lines, i-1, j-1) == "M" &&
					getChar(lines, i+1, j+1) == "S" &&
					getChar(lines, i-1, j+1) == "S" &&
					getChar(lines, i+1, j-1) == "M" {
					sum++
				}

				//S.M
				//.A.
				//S.M
				if getChar(lines, i-1, j-1) == "S" &&
					getChar(lines, i+1, j+1) == "M" &&
					getChar(lines, i-1, j+1) == "S" &&
					getChar(lines, i+1, j-1) == "M" {
					sum++
				}

				//S.S
				//.A.
				//M.M
				if getChar(lines, i-1, j-1) == "S" &&
					getChar(lines, i+1, j+1) == "M" &&
					getChar(lines, i-1, j+1) == "M" &&
					getChar(lines, i+1, j-1) == "S" {
					sum++
				}
			}
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
