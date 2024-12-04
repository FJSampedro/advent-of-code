package main

import (
	"fmt"
	"lib"
	"strconv"
	"strings"
)

func readInputTxt() (array1 []int, array2 []int) {

	lines, err := lib.ReadTxtLines()

	if err != nil {
		return
	}

	// Read lines
	for _, line := range lines {
		// Strips in columns
		columns := strings.Fields(line)
		if len(columns) != 2 {
			fmt.Println("Format not valid in line:", line)
			continue
		}

		// Str to int conversion
		num1, err1 := strconv.Atoi(columns[0])
		num2, err2 := strconv.Atoi(columns[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error in string conversion at line:", line)
			continue
		}

		// Adds to arrays
		array1 = append(array1, num1)
		array2 = append(array2, num2)
	}

	return
}

func getTotalDistance(array1 []int, array2 []int) (total int) {
	lib.QuickSort(array1)

	lib.QuickSort(array2)

	for i := range array1 {

		if distance := array2[i] - array1[i]; distance > 0 {
			total += distance
		} else {
			distance := array1[i] - array2[i]
			total += distance
		}
	}
	return total
}

func getSimilarityScore(array1 []int, array2 []int) (similarity int) {

	for _, num1 := range array1 {
		var mult int
		for _, num2 := range array2 {
			if num1 == num2 {
				mult++
			}
		}
		similarity += num1 * mult
	}

	return similarity
}

func main() {

	fmt.Println("############# Part One ##############")

	arr1, arr2 := readInputTxt()

	sum := getTotalDistance(arr1, arr2)

	fmt.Println("Total distance : ", sum)

	fmt.Println("############# Part Two ##############")

	arr1, arr2 = readInputTxt()

	sim := getSimilarityScore(arr1, arr2)

	fmt.Println("Similarity Score : ", sim)
}
