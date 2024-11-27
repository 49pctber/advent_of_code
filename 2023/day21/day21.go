package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	up int = iota
	down
	left
	right
)

type plot_t struct {
	category  byte
	processed bool
	reachable bool
	dist      int // distance from starting point
}

type search_t struct {
	grid [][]plot_t
	row  int
	col  int
	n    int
	dist int
}

func (search *search_t) Step(direction int) *search_t {
	new_search := search_t{
		grid: search.grid,
		row:  search.row,
		col:  search.col,
		n:    search.n - 1,
		dist: search.dist + 1,
	}

	switch direction {
	case up:
		new_search.row--
	case down:
		new_search.row++
	case right:
		new_search.col++
	case left:
		new_search.col--
	}

	// new_search.row = (new_search.row + len(search.grid)) % len(search.grid)
	// new_search.col = (new_search.col + len(search.grid[0])) % len(search.grid[0])

	return &new_search
}

func (search search_t) String() string {
	rows := []string{}
	for i := range search.grid {
		var row []byte
		for _, plot := range search.grid[i] {
			if plot.reachable {
				row = append(row, 'O')
			} else {
				row = append(row, plot.category)
			}
		}
		rows = append(rows, string(row))
	}

	return strings.Join(rows, "\n")
}

func (search search_t) PrintDist() {
	rows := []string{}
	for i := range search.grid {
		var row []string
		if len(search.grid[i]) != 131 {
			log.Fatalf("Not enough columns %d", len(search.grid[i]))
		}
		for j := range search.grid[i] {
			if search.grid[i][j].dist != math.MaxInt {
				row = append(row, strconv.Itoa(search.grid[i][j].dist))
			} else {
				row = append(row, "#")
			}
		}
		rows = append(rows, strings.Join(row, "\t"))
	}

	fmt.Println(strings.Join(rows, "\n"))
}

var ErrStartingPointNotFound error = errors.New("starting point not found")

func (search search_t) FindStart() (int, int, error) {
	for i := range search.grid {
		for j := range search.grid[i] {
			if search.grid[i][j].category == 'S' {
				return i, j, nil
			}
		}
	}
	return 0, 0, ErrStartingPointNotFound
}

func (search search_t) CountReachable() int {
	count := 0
	for i := range search.grid {
		for j := range search.grid[i] {
			if search.grid[i][j].reachable {
				count++
			}
		}
	}
	return count
}

func FillIn(search *search_t) []*search_t {

	// check if out of bounds
	if search.row < 0 || search.col < 0 || search.row >= len(search.grid) || search.col >= len(search.grid[search.row]) {
		return nil
	}

	// check for rocks
	if search.grid[search.row][search.col].category == '#' {
		return nil
	}

	// check if already max_remaining
	if search.grid[search.row][search.col].processed {
		return nil
	} else {
		search.grid[search.row][search.col].dist = search.dist
		search.grid[search.row][search.col].processed = true
	}

	// check if reachable
	if search.n%2 == 0 {
		search.grid[search.row][search.col].reachable = true
	}

	// step to adjacent plots if n > 0
	if search.n > 0 {
		next_steps := []*search_t{}

		for _, next_search := range []*search_t{search.Step(up), search.Step(down), search.Step(right), search.Step(left)} {
			// if !search.grid[next_search.row][next_search.col].processed {
			next_steps = append(next_steps, next_search)
			// }
		}

		return next_steps
	}

	return nil
}

func CountGardenPlots(s string, n int) int {
	search := ParseDay21Input(s, n)

	var err error

	search.row, search.col, err = search.FindStart()
	if err != nil {
		log.Fatal(err)
	}
	search.n = n

	queue := []*search_t{&search}
	for len(queue) > 0 {
		var s search_t
		s, queue = *queue[0], queue[1:]
		queue = append(queue, FillIn(&s)...)
	}

	return search.CountReachable()
}

func ParseDay21Input(s string, n int) search_t {
	var search search_t

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]plot_t, len(scanner.Text()))
		for i, c := range scanner.Text() {
			row[i].category = byte(c)
			row[i].dist = math.MaxInt
		}
		search.grid = append(search.grid, row)
	}

	search.n = n

	search.row, search.col, err = search.FindStart()
	if err != nil {
		log.Fatal(err)
	}

	return search
}

func CountGardenPlots2(s string, n int) int {
	c_ef := CountGardenPlots(s, 132)
	c_of := CountGardenPlots(s, 131)
	c_ec := c_ef - CountGardenPlots(s, 64)
	c_oc := c_of - CountGardenPlots(s, 65)

	q := n / 131
	r := n % 131

	if q != 202300 {
		panic("q should be 202300")
	}

	if q%2 != 0 {
		panic("q must be even")
	}

	if r != 65 {
		panic("formula only valid when r = 65")
	}

	return (q+1)*((q+1)*c_of-c_oc) + q*(q*c_ef+c_ec)
}

func main() {
	fmt.Println("Day 21")

	part1 := CountGardenPlots(`input21.txt`, 64)
	fmt.Printf("part1: %v\n", part1)

	part2 := CountGardenPlots2(`input21.txt`, 26501365)
	fmt.Printf("part2: %v\n", part2)
}
