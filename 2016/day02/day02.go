package main

import (
	"fmt"
	"strings"

	_ "embed"
)

func GetCode(input string) string {

	var code string = ""

	keypad := make([][]int, 3)
	for i := range keypad {
		keypad[i] = make([]int, 3)
	}

	i := 1
	for row := range []int{0, 1, 2} {
		for col := range []int{0, 1, 2} {
			keypad[row][col] = i
			i++
		}
	}

	x := 1
	y := 1

	instructions := strings.Split(input, "\n")

	for _, instruction := range instructions {
		if len(instruction) == 0 {
			continue
		}

		for _, b := range instruction {
			switch b {
			case 'U':
				y -= 1
				if y < 0 {
					y = 0
				}
			case 'R':
				x += 1
				if x > 2 {
					x = 2
				}
			case 'L':
				x -= 1
				if x < 0 {
					x = 0
				}
			case 'D':
				y += 1
				if y > 2 {
					y = 2
				}
			}
		}

		code = code + fmt.Sprintf("%d", keypad[y][x])

	}

	return code
}

func checkBounds(row, col int) bool {
	a, b := row-2, col-2
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	sum := a + b
	return sum <= 2
}

func GetCode2(input string) string {

	var code string = ""

	keypad := make([][]int, 5)
	for i := range keypad {
		keypad[i] = make([]int, 5)
	}

	i := 1
	for row := range []int{0, 1, 2, 3, 4} {
		for col := range []int{0, 1, 2, 3, 4} {
			if checkBounds(row, col) {
				keypad[row][col] = i
				i++
			}
		}
	}

	x := 0
	y := 2

	instructions := strings.Split(input, "\n")

	for _, instruction := range instructions {
		if len(instruction) == 0 {
			continue
		}

		for _, b := range instruction {
			switch b {
			case 'U':
				if checkBounds(y-1, x) {
					y -= 1
				}
			case 'R':
				if checkBounds(y, x+1) {
					x += 1
				}
			case 'L':
				if checkBounds(y, x-1) {
					x -= 1
				}
			case 'D':
				if checkBounds(y+1, x) {
					y += 1
				}
			}
		}
		code = code + fmt.Sprintf("%X", keypad[y][x])
	}

	return code
}

//go:embed input/input
var input string

func main() {
	fmt.Printf("Part 1: %s\n", GetCode(input))
	fmt.Printf("Part 2: %s\n", GetCode2(input))
}
