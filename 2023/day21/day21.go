/*
This code is not in a functional state.
*/
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

func getDistanceGrid(s string, row, col, n int) [][]plot_t {
	search := ParseDay21Input(s, n) // TODO maybe n+1?
	search.row = row
	search.col = col

	queue := []*search_t{&search}
	for len(queue) > 0 {
		var s search_t
		s, queue = *queue[0], queue[1:]
		queue = append(queue, FillIn(&s)...)
	}

	// fmt.Println(search.CountReachable())
	// search.PrintDist()

	return search.grid
}

func CountGardenPlots2(s string, n int) int {
	search := ParseDay21Input(s, n)

	width := len(search.grid)
	height := len(search.grid[0])

	c0 := CountGardenPlots(s, n)
	c1 := CountGardenPlots(s, n+1)
	fmt.Println("c0 c1:", c0, c1)

	q, r := n/width, n%width
	if q != 202300 || r != 65 {
		log.Fatalf("q,r are wrong (%d,%d)", q, r)
	}

	N := (q - 1) / 2
	M := q / 2
	enclosed := c0*(4*N*(N+1)+1) + c1*(4*M*M) // completely enclosed
	if enclosed != 625621818307437 {
		log.Fatalf("enclosed is incorrect (%d)", enclosed)
	}

	pn := getDistanceGrid(s, 0, (width-1)/2, 100000+q)
	pe := getDistanceGrid(s, (height-1)/2, width-1, 100000+q)
	ps := getDistanceGrid(s, height-1, (width-1)/2, 100000+q)
	pw := getDistanceGrid(s, (height-1)/2, 0, 100000+q)

	pnw := getDistanceGrid(s, 0, 0, 100000+q)
	pne := getDistanceGrid(s, 0, width-1, 100000+q)
	psw := getDistanceGrid(s, height-1, 0, 100000+q)
	pse := getDistanceGrid(s, height-1, width-1, 100000+q)

	pnwi := getDistanceGrid(s, 0, 0, 100000+q+1)
	pnei := getDistanceGrid(s, 0, width-1, 100000+q+1)
	pswi := getDistanceGrid(s, height-1, 0, 100000+q+1)
	psei := getDistanceGrid(s, height-1, width-1, 100000+q+1)

	edge_count := 0
	for i := 0; i < height-1; i++ {
		for j := 0; j < width-1; j++ {

			if pn[i][j].reachable && pn[i][j].dist <= 130 {
				edge_count += 1
			}
			if pe[i][j].reachable && pe[i][j].dist <= 130 {
				edge_count += 1
			}
			if ps[i][j].reachable && ps[i][j].dist <= 130 {
				edge_count += 1
			}
			if pw[i][j].reachable && pw[i][j].dist <= 130 {
				edge_count += 1
			}

			if pse[i][j].reachable && pse[i][j].dist <= 130+65 {
				edge_count += q - 1
			}
			if pne[i][j].reachable && pne[i][j].dist <= 130+65 {
				edge_count += q - 1
			}
			if psw[i][j].reachable && psw[i][j].dist <= 130+65 {
				edge_count += q - 1
			}
			if pnw[i][j].reachable && pnw[i][j].dist <= 130+65 {
				edge_count += q - 1
			}

			if psei[i][j].reachable && psei[i][j].dist < 65 {
				edge_count += q
			}
			if pnei[i][j].reachable && pnei[i][j].dist < 65 {
				edge_count += q
			}
			if pswi[i][j].reachable && pswi[i][j].dist < 65 {
				edge_count += q
			}
			if pnwi[i][j].reachable && pnwi[i][j].dist < 65 {
				edge_count += q
			}
		}
	}

	return enclosed + edge_count
}

func main() {
	part1 := CountGardenPlots(`input\input21.txt`, 64)
	fmt.Printf("part1: %v\n", part1) // 3773

	part2 := CountGardenPlots2(`input\input21.txt`, 26501365)
	fmt.Printf("part2: %v\n", part2) // 625628021226274

	if part2 != 625628021226274 {
		diff := 625628021226274 - part2
		q := diff / 131
		r := diff % 131
		log.Fatalf("part2 should be 625628021226274 (diff: %d) [%d, %d]", diff, q, r)
	}
}
