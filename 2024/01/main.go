package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read_input_txt() (array1 []int, array2 []int) {

	// Open file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

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

	// Error handling during reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	return
}

func get_total_distance(array1 []int, array2 []int) (total int) {

	sort.Ints(array1)
	sort.Ints(array2)

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

func get_similarity_score(array1 []int, array2 []int) (similarity int) {

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
	arr1, arr2 := read_input_txt()

	sum := get_total_distance(arr1, arr2)

	fmt.Println("Total distance : ", sum)

	fmt.Println("############# Part Two ##############")

	arr1, arr2 = read_input_txt()

	sim := get_similarity_score(arr1, arr2)

	fmt.Println("Similarity Score : ", sim)
}
