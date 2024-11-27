package main

import (
	"bufio"
	"fmt"
	"os"
)

var state [2][100][100]byte
var count int

func printState() {
	fmt.Printf("[count: %d]\n", count)
	for _, row := range state[count%2] {
		fmt.Printf("%s\n", row)
	}
	fmt.Printf("[count: %d]\n\n", count)
}

func countneighbors(x, y int) int {
	n := 0
	for dx := -1; dx <= 1; dx++ {
		if dx+x == -1 || dx+x == 100 {
			continue
		}
		for dy := -1; dy <= 1; dy++ {
			if dy+y == -1 || dy+y == 100 {
				continue
			} else if dx == dy && dx == 0 {
				continue
			}
			if state[count%2][x+dx][y+dy] == '#' {
				n++
			}
		}
	}
	return n
}

func countOn() int {
	c := 0
	for _, row := range state[count%2] {
		for _, b := range row {
			if b == '#' {
				c++
			}
		}
	}
	return c
}

func stepV1() {
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			c := countneighbors(x, y)
			lightstate := state[count%2][x][y]

			if lightstate == '#' {
				if c == 2 || c == 3 {
					state[(count+1)%2][x][y] = '#'
				} else {
					state[(count+1)%2][x][y] = '.'
				}
			} else {
				if c == 3 {
					state[(count+1)%2][x][y] = '#'
				} else {
					state[(count+1)%2][x][y] = '.'
				}
			}

		}
	}
	count++
}

func stepV2() {

	state[count%2][0][0] = '#'
	state[count%2][0][99] = '#'
	state[count%2][99][0] = '#'
	state[count%2][99][99] = '#'
	state[(count+1)%2][0][0] = '#'
	state[(count+1)%2][0][99] = '#'
	state[(count+1)%2][99][0] = '#'
	state[(count+1)%2][99][99] = '#'

	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			c := countneighbors(x, y)
			lightstate := state[count%2][x][y]

			if lightstate == '#' {
				if c == 2 || c == 3 {
					state[(count+1)%2][x][y] = '#'
				} else {
					state[(count+1)%2][x][y] = '.'
				}
			} else {
				if c == 3 {
					state[(count+1)%2][x][y] = '#'
				} else {
					state[(count+1)%2][x][y] = '.'
				}
			}

		}
	}

	count++

	state[count%2][0][0] = '#'
	state[count%2][0][99] = '#'
	state[count%2][99][0] = '#'
	state[count%2][99][99] = '#'
	state[(count+1)%2][0][0] = '#'
	state[(count+1)%2][0][99] = '#'
	state[(count+1)%2][99][0] = '#'
	state[(count+1)%2][99][99] = '#'
}

func part1() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)

	for i := 0; scanner.Scan(); i++ {
		copy(state[0][i][:], scanner.Bytes())
	}

	count = 0
	for i := 0; i < 100; i++ {
		stepV1()
	}

	fmt.Println("Part 1:", countOn())
}

func part2() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)

	for i := 0; scanner.Scan(); i++ {
		copy(state[0][i][:], scanner.Bytes())
	}

	count = 0
	for i := 0; i < 100; i++ {
		stepV2()
	}

	fmt.Println("Part 2:", countOn())
}

func main() {
	fmt.Println("Day 18")
	part1()
	part2()
}
