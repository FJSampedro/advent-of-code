package main

import (
	"fmt"
	"lib"
	"strconv"
	"strings"
)

func getOrdersNRules() (map[int][]int, [][]int) {
	lines, err := lib.ReadTxtLines()
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	list_orders := false
	var stringRules []string
	var stringOrders []string
	for _, line := range lines {
		if line != "" {
			if list_orders {
				stringOrders = append(stringOrders, line)
			} else {
				stringRules = append(stringRules, line)
			}
		} else {
			list_orders = true
		}
	}

	rules := make(map[int][]int)
	for _, r := range stringRules {
		parts := strings.Split(r, "|")
		key, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting key:", err)
		}
		value, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting value:", err)
		}
		if _, exists := rules[key]; exists {
			rules[key] = append(rules[key], value)
		} else {
			rules[key] = []int{value}
		}
	}

	var orders [][]int
	for _, o := range stringOrders {
		parts := strings.Split(o, ",")
		var order []int
		for _, page := range parts {
			p, err := strconv.Atoi(page)
			if err != nil {
				fmt.Println("Error converting key:", err)
			}
			order = append(order, p)
		}
		orders = append(orders, order)
	}
	return rules, orders
}
func getCorrectOrders(rules map[int][]int, orders [][]int) (correctOrders []int) {
	for orderNumber, order := range orders {
		correct := true
		for i, page := range order {
			if rule, exists := rules[page]; exists {
				for _, n := range order[i+1:] {
					for _, r := range rule {
						if n == r {
							correct = false
						}

					}
				}
			}
		}
		if correct {
			correctOrders = append(correctOrders, orderNumber)
		}
	}
	return
}
func getMiddlePage(order []int) (page int) {
	page = order[len(order)/2]
	return
}
func partOne() {
	rules, orders := getOrdersNRules()
	correct := getCorrectOrders(rules, orders)
	sum := 0
	for _, c := range correct {
		sum += getMiddlePage(orders[c])
	}
	fmt.Println(sum)
}

func getIncorrectOrders(rules map[int][]int, orders [][]int) (incorrectOrders []int) {
	for orderNumber, order := range orders {
		incorrect := false
		for i, page := range order {
			if rule, exists := rules[page]; exists {
				for _, n := range order[i+1:] {
					for _, r := range rule {
						if n == r {
							incorrect = true
						}

					}
				}
			}
		}
		if incorrect {
			incorrectOrders = append(incorrectOrders, orderNumber)
		}
	}
	return
}

func fixOrders(rules map[int][]int, order []int) (incorrect bool) {
	incorrect = false
	for i, page := range order {
		if rule, exists := rules[page]; exists {
			for j, n := range order[i+1:] {
				for _, r := range rule {
					if n == r {
						incorrect = true
						temp := order[i]
						order[i] = order[j+i+1]
						order[j+i+1] = temp
					}
				}
			}
		}
	}
	if incorrect {
		fixOrders(rules, order)
	}
	return incorrect
}

func partTwo() {
	rules, orders := getOrdersNRules()
	incorrect := getIncorrectOrders(rules, orders)

	for _, o := range incorrect {
		fixOrders(rules, orders[o])
	}
	sum := 0
	for _, c := range incorrect {
		sum += getMiddlePage(orders[c])
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println("################## Part One ##################")
	partOne()
	fmt.Println("################## Part Two ##################")
	partTwo()
}
