package main

import (
	"fmt"
	"lib"
	"strconv"
	"strings"
)

func readInputTxt() (array [][]int) {
	lines, err := lib.ReadTxtLines()
	if err != nil {
		return
	}

	// Read lines
	for _, line := range lines {
		var aux_array []int
		// Strips in columns
		svalues := strings.Fields(line)
		// Str to int conversion

		for i, svalue := range svalues {
			value, err1 := strconv.Atoi(svalue)
			if err1 != nil {
				fmt.Println("Error in string conversion at line:", line, " index:", i)
				continue
			}
			aux_array = append(aux_array, value)
		}
		array = append(array, aux_array)
	}

	return
}
func partOne(inputs [][]int) {
	reports_sum := 0
	for _, report := range inputs {
		safe := true
		var increasing bool
		for i := range report[1:] {
			if i == 0 {
				if report[0] < report[1] {
					increasing = true
				} else {
					increasing = false
				}
			}
			if increasing {
				if report[i+1] > report[i]+3 {
					safe = false
				}
				if report[i+1] < report[i] {
					safe = false
				}
			} else {
				if report[i] > report[i+1]+3 {
					safe = false
				}
				if report[i] < report[i+1] {
					safe = false
				}
			}
			if report[i] == report[i+1] {
				safe = false
			}

		}
		if safe {
			reports_sum++
		}
	}
	fmt.Println("The num of safe reports are : ", reports_sum)
}

func checkSafe(report []int) (safe int) {
	safe = 0
	increasing := false
	decreasing := false

	for i := range report[1:] {
		if i == 0 {
			if report[0] < report[1] {
				increasing = true
			} else if report[0] > report[1] {
				decreasing = true
			} else {
				safe = i + 1
			}
		}
		if increasing {
			if report[i+1] > report[i]+3 {
				safe = i + 1
			}
			if report[i+1] < report[i] {
				safe = i + 1
			}
		} else if decreasing {
			if report[i] > report[i+1]+3 {
				safe = i + 1
			}
			if report[i] < report[i+1] {
				safe = i + 1
			}
		}
		if report[i] == report[i+1] {
			safe = i + 1
		}

	}
	return
}

func partTwo(inputs [][]int) {
	reports_sum := 0
	for _, report := range inputs {
		safe := checkSafe(report)
		if safe != 0 {
			//fmt.Println(report)
			for i := range report {
				var check_subs []int
				check_subs = append(check_subs, report[:i]...)
				check_subs = append(check_subs, report[i+1:]...)
				safe = checkSafe(check_subs)
				if safe == 0 {
					break
				}
			}
		}
		if safe == 0 {
			reports_sum++
		}

	}
	fmt.Println("The num of safe reports are : ", reports_sum)
}

func main() {
	fmt.Println("################## Part One ##################")
	inputs := readInputTxt()
	partOne(inputs)
	fmt.Println("################## Part Two ##################")
	inputs = readInputTxt()
	partTwo(inputs)
}
