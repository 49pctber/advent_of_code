package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Day 1")

	input, err := os.Open(`input\input1.txt`)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Scan()

	floor := 0
	basement := 0
	for i, c := range scanner.Text() {
		if c == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 && basement == 0 {
			basement = i + 1
		}
	}

	fmt.Println("Floor: ", floor)             // 232
	fmt.Println("Basement Index: ", basement) // 1783
}
