package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type location struct {
	x int
	y int
}

func (l location) Right() location {
	return location{x: l.x + 1, y: l.y}
}

func (l location) Left() location {
	return location{x: l.x - 1, y: l.y}
}

func (l location) Up() location {
	return location{x: l.x, y: l.y - 1}
}

func (l location) Down() location {
	return location{x: l.x, y: l.y + 1}
}

var grid [140][140]rune
var distance [140][140]int

var big_grid [279][279]rune

func assignDistance(l location, d int) {
	if distance[l.y][l.x] != -1 && d >= distance[l.y][l.x] {
		return
	}

	distance[l.y][l.x] = d

	switch grid[l.y][l.x] {
	case '-':
		assignDistance(l.Right(), d+1)
		assignDistance(l.Left(), d+1)
	case '|':
		assignDistance(l.Up(), d+1)
		assignDistance(l.Down(), d+1)
	case 'J':
		assignDistance(l.Up(), d+1)
		assignDistance(l.Left(), d+1)
	case 'L':
		assignDistance(l.Up(), d+1)
		assignDistance(l.Right(), d+1)
	case '7':
		assignDistance(l.Down(), d+1)
		assignDistance(l.Left(), d+1)
	case 'F':
		assignDistance(l.Down(), d+1)
		assignDistance(l.Right(), d+1)
	}
}

func markOutside(l location) {
	if l.x < 0 || l.x >= 279 {
		return
	}

	if l.y < 0 || l.y >= 279 {
		return
	}

	if big_grid[l.y][l.x] == '.' {
		big_grid[l.y][l.x] = 'O'
		markOutside(l.Down())
		markOutside(l.Up())
		markOutside(l.Right())
		markOutside(l.Left())
	}
}

func setBigGrid(l location, r rune) {
	if l.x < 0 || l.x >= 279 {
		return
	}

	if l.y < 0 || l.y >= 279 {
		return
	}

	big_grid[l.y][l.x] = r
}

func main() {
	fmt.Println("Day 10")

	file, err := os.Open(filepath.Join("input", "input10.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for row := 0; row < 140; row++ {
		for col := 0; col < 140; col++ {
			distance[row][col] = -1
		}
	}

	row := 0
	var start location

	for scanner.Scan() {
		for col, b := range scanner.Text() {
			grid[row][col] = b
			if b == 'S' {
				start = location{x: col, y: row}
				grid[row][col] = 'J' // hacky solution
				fmt.Println(start)
			}
		}
		row++
	}

	// find rest of distances
	assignDistance(start, 0)

	max := 0

	for _, row := range distance {
		for _, val := range row {
			if val > max {
				max = val
			}
		}
	}

	fmt.Printf("max: %v\n", max) // 7030

	// Part 2

	for row := 0; row < 279; row++ {
		for col := 0; col < 279; col++ {
			big_grid[row][col] = '.'
		}
	}

	for row := 0; row < 140; row++ {
		for col := 0; col < 140; col++ {
			if distance[row][col] == -1 {
				continue
			}
			l := location{x: 2 * col, y: 2 * row}
			setBigGrid(l, grid[row][col])

			switch grid[row][col] {
			case '-':
				setBigGrid(l.Right(), '-')
				setBigGrid(l.Left(), '-')
			case '|':
				setBigGrid(l.Up(), '|')
				setBigGrid(l.Down(), '|')
			case 'J':
				setBigGrid(l.Left(), '-')
				setBigGrid(l.Up(), '|')
			case 'L':
				setBigGrid(l.Right(), '-')
				setBigGrid(l.Up(), '|')
			case '7':
				setBigGrid(l.Left(), '-')
				setBigGrid(l.Down(), '|')
			case 'F':
				setBigGrid(l.Right(), '-')
				setBigGrid(l.Down(), '|')
			}
		}
	}

	markOutside(location{x: 0, y: 0})

	count := 0

	for row := 0; row < 140; row++ {
		for col := 0; col < 140; col++ {
			if big_grid[2*row][2*col] == '.' {
				count++
			}
		}
	}

	fmt.Printf("count: %v\n", count) // 285
}
