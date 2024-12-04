package lib

import (
	"bufio"
	"fmt"
	"os"
)

func ReadTxtLines() (lines []string, err error) {

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
		// Adds to arrays
		lines = append(lines, line)
	}

	// Error handling during reading
	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	return
}
