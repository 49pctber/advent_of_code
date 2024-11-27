package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3_1() {
	input, err := os.ReadFile("input/input3.txt")

	rows := strings.Split(string(input), "\n")

	if err != nil {
		panic(err)
	}

	var v [140][140]bool

	for i, row := range rows {
		for j, col := range row {
			if !strings.Contains(".1234567890", string(col)) {
				for _, m := range []int{i - 1, i, i + 1} {
					for _, n := range []int{j - 1, j, j + 1} {
						v[m][n] = true
					}
				}
			}
		}
	}

	sum := 0
	re := regexp.MustCompile(`(\d)+`)

	for i, row := range rows {
		matches := re.FindAllStringSubmatchIndex(row, -1)
		for _, j := range matches {
			// fmt.Println(row[j[0]:j[1]])
			add_to_sum := false

			for _, x := range v[i][j[0]:j[1]] {
				if x {
					add_to_sum = true
					break
				}
			}

			if add_to_sum {
				n, err := strconv.Atoi(row[j[0]:j[1]])
				if err != nil {
					panic(err)
				}
				sum += n
			}
		}
	}

	fmt.Println("Sum:", sum) // Answer: 525181
}

func Day3_2() {

	input, err := os.ReadFile("input/input3.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(input), "\n")

	// array for storing products of adjacent numbers
	var prod [140][140]uint64

	for i := 0; i < 140; i++ {
		for j := 0; j < 140; j++ {
			if rows[i][j] == '*' {
				prod[i][j] = 1
			}
		}
	}

	// create array of ints to keep track of number of adjacent numbers
	var adj [140][140]int

	// find all numbers, increment adjacent cells
	re := regexp.MustCompile(`(\d)+`)
	for i, row := range rows {
		matches := re.FindAllStringSubmatchIndex(row, -1)
		for _, j := range matches {

			n, err := strconv.Atoi(row[j[0]:j[1]])
			if err != nil {
				panic(err)
			}

			jlow := j[0] - 1
			jhigh := j[1] + 1

			if jlow < 0 {
				jlow = 0
			}
			if jhigh >= 140 {
				jhigh = 139
			}

			ilow := i - 1
			ihigh := i + 1

			if ilow < 0 {
				ilow = 0
			}
			if ihigh >= 140 {
				ihigh = 139
			}
			for i := ilow; i <= ihigh; i++ {
				for j := jlow; j < jhigh; j++ {
					adj[i][j]++
					prod[i][j] *= uint64(n)
				}
			}
		}
	}

	// check for cells containing *
	var sum uint64 = 0

	for i := 0; i < 140; i++ {
		for j := 0; j < 140; j++ {
			if adj[i][j] == 2 {
				sum += prod[i][j]
			}
		}
	}

	fmt.Println("Sum:", sum) // Answer: 84289137

}

func main() {
	fmt.Println("Day 3")
	Day3_1()
	Day3_2()
}
