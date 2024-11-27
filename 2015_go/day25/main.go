package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func indexAt(row, col int) int {
	if row < 1 || col < 1 {
		panic("invalid row/column")
	}

	n := 1 // number of iterations
	for sum := 2; sum < row+col; sum++ {
		n += sum
	}
	n -= row - 1

	return n
}

func valAt(row, col int) int64 {
	var val int64 = 20151125

	n := indexAt(row, col) - 1 // number of iterations

	for ; n > 0; n-- {
		val *= 252533
		val %= 33554393
	}

	return val
}

func main() {

	fmt.Println("Day 25")

	re := regexp.MustCompile(`To continue, please consult the code grid in the manual.  Enter the code at row (\d+), column (\d+).`)
	input, err := os.ReadFile(`input.txt`)
	if err != nil {
		panic(err)
	}

	m := re.FindSubmatch(input)

	row, _ := strconv.Atoi(string(m[1]))
	col, _ := strconv.Atoi(string(m[2]))

	fmt.Println("Part 1:", valAt(row, col))
}
