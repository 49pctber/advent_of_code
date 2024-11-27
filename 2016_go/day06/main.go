package main

import (
	"fmt"
	"strings"

	_ "embed"
)

//go:embed input
var input string

func Decode(input string) string {
	rows := strings.Split(input, "\n")

	// initialize maps
	freq := make([]map[byte]int, len(rows[0]))
	for i := range freq {
		freq[i] = make(map[byte]int, 26)
	}

	// count letter frequency
	for _, row := range rows {
		for i := 0; i < len(row); i++ {
			freq[i][row[i]]++
		}
	}

	// find most common letters
	most_common := make([]byte, len(rows[0]))
	for i := range freq {
		m := 0
		for key, value := range freq[i] {
			if value > m {
				m = value
				most_common[i] = key
			}
		}
	}

	return string(most_common)
}

func Decode2(input string) string {
	rows := strings.Split(input, "\n")

	// initialize maps
	freq := make([]map[byte]int, len(rows[0]))
	for i := range freq {
		freq[i] = make(map[byte]int, 26)
	}

	// count letter frequency
	for _, row := range rows {
		for i := 0; i < len(row); i++ {
			freq[i][row[i]]++
		}
	}

	// find least common letters
	least_common := make([]byte, len(rows[0]))
	for i := range freq {
		m := 10000000
		for key, value := range freq[i] {
			if value < m {
				m = value
				least_common[i] = key
			}
		}
	}

	return string(least_common)
}

func Part1() {
	fmt.Printf("Part 1: %v\n", Decode(input))
}

func Part2() {
	fmt.Printf("Part 2: %v\n", Decode2(input))
}

func main() {
	Part1()
	Part2()
}
