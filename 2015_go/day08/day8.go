package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Day 8")

	file, err := os.Open(`input.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// load instructions
	literal_sum := 0
	memory_sum := 0
	encoded_sum := 0
	for scanner.Scan() {
		literal_sum += len(scanner.Text())
		for i := 1; i < len(scanner.Text())-1; i++ {
			if scanner.Text()[i] == '\\' {
				if scanner.Text()[i+1] == 'x' {
					i += 3
				} else {
					i += 1
				}
			}
			memory_sum++
		}

		encoded_sum += 2 // adding two quote characters
		encoded_sum += len(scanner.Text())
		for i := 0; i < len(scanner.Text()); i++ {
			if scanner.Text()[i] == '\\' || scanner.Text()[i] == '"' {
				encoded_sum++
			}
		}

	}

	diff := literal_sum - memory_sum
	fmt.Printf("diff: %v\n", diff) // 1333

	diff2 := encoded_sum - literal_sum
	fmt.Printf("diff: %v\n", diff2) // 2046
}
