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

	// Abre el archivo
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Lector de líneas
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Divide la línea en columnas
		columns := strings.Fields(line)
		if len(columns) != 2 {
			fmt.Println("Formato inválido en la línea:", line)
			continue
		}

		// Convierte las columnas a enteros
		num1, err1 := strconv.Atoi(columns[0])
		num2, err2 := strconv.Atoi(columns[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error al convertir números en la línea:", line)
			continue
		}

		// Agrega a los arrays
		array1 = append(array1, num1)
		array2 = append(array2, num2)
	}

	// Manejo de errores durante la lectura
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	return
}

func main() {

	fmt.Println("############# Part One ##############")
	arr1, arr2 := read_input_txt()

	sort.Ints(arr1)
	sort.Ints(arr2)

	var sum int
	for i, _ := range arr1 {

		if distance := arr2[i] - arr1[i]; distance > 0 {
			sum += distance
		} else {
			distance := arr1[i] - arr2[i]
			sum += distance
		}
	}
	fmt.Println("Total distance : ", sum)

	fmt.Println("############# Part Two ##############")

	arr1, arr2 = read_input_txt()

	var sim int
	for _, num1 := range arr1 {
		var mult int
		for _, num2 := range arr2 {
			if num1 == num2 {
				mult++
			}
		}
		sim += num1 * mult
	}

	fmt.Println("Similarity Score : ", sim)
}
