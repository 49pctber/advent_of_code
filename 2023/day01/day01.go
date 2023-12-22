package main

import (
	"fmt"
	"os"
	"strings"
)

var lut = make(map[string]int)

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func parseString(input string) int {

	var first bool = true
	var last_digit int = 0
	var first_digit int = 0

	for _, char := range input {
		if isDigit(byte(char)) {
			last_digit = int(char - '0')
			if first {
				first_digit = last_digit
				first = false
			}
		}
	}

	return 10*first_digit + last_digit
}

func convertToDigits(input string) string {
	var output string = ""
	for i := range input {
		if isDigit(input[i]) {
			output += string(input[i])
			continue
		}

		for _, j := range []int{3, 4, 5} {
			if i <= len(input)-j {
				val, ok := lut[input[i:i+j]]
				if ok {
					output = fmt.Sprintf("%s%d", output, val)
					continue
				}
			}
		}

	}

	return output
}

func challenge2() {
	fmt.Println("Challenge 2")

	input, err := os.ReadFile("input/input1.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(input), "\n")

	total := 0

	for _, row := range rows {
		total += parseString(convertToDigits(row))
	}

	fmt.Println(total)
}

func challenge1() {
	// Problem 1
	fmt.Println("Challenge 1")

	input, err := os.ReadFile("input/input1.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(input), "\n")

	total := 0

	for _, row := range rows {
		total += parseString(row)
	}

	fmt.Println(total)
}

func init() {
	lut["one"] = 1
	lut["two"] = 2
	lut["three"] = 3
	lut["four"] = 4
	lut["five"] = 5
	lut["six"] = 6
	lut["seven"] = 7
	lut["eight"] = 8
	lut["nine"] = 9
}

func main() {
	fmt.Println("Day 1")
	challenge1() // 54390
	challenge2() // 54277
}
