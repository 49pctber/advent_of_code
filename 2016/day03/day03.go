package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input/input
var input string

func checkTriangle(a, b, c int) bool {
	sum := a + b + c
	m := max(a, b, c)
	return sum-2*m > 0
}

func Part1() {
	nvalid := 0
	rows := strings.Split(input, "\n")

	for _, row := range rows {
		if len(row) < 3 {
			continue
		}

		var a, b, c int
		fmt.Sscanf(row, " %d %d %d", &a, &b, &c)

		if checkTriangle(a, b, c) {
			nvalid++
		}

	}

	fmt.Printf("Part 1: %v\n", nvalid)
}

func Part2() {
	nvalid := 0
	rows := strings.Split(input, "\n")
	x := make([]int, 9)
	count := 0

	for _, row := range rows {
		if len(row) < 3 {
			continue
		}

		fmt.Sscanf(row, " %d %d %d", &x[count], &x[count+1], &x[count+2])
		count += 3

		if count == 9 {
			count = 0

			if checkTriangle(x[0], x[3], x[6]) {
				nvalid++
			}
			if checkTriangle(x[1], x[4], x[7]) {
				nvalid++
			}
			if checkTriangle(x[2], x[5], x[8]) {
				nvalid++
			}
		}
	}
	fmt.Printf("Part 2: %v\n", nvalid)
}

func main() {
	Part1()
	Part2()
}
