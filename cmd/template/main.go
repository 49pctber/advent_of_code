package main

import (
	"fmt"
	"os"

	_ "embed"
)

func Part1() {
	fmt.Printf("Part 1: %v\n", "?")
}

func Part2() {
	fmt.Printf("Part 2: %v\n", "?")
}

func InputString(fname string) string {
	b, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	Part1()
	Part2()
}
