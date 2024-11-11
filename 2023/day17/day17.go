/*
This has a bug somewhere in it.
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

const (
	up int = iota
	down
	left
	right
)

type step_t struct {
	x         int
	y         int
	heat_loss int
	count     int
	dir       int
}

var ErrInvalidStep error = errors.New("invalid step")
var best int
var show_best bool = false
var min_dir int
var max_dir int

func (s step_t) Duplicate() step_t {
	return step_t{x: s.x, y: s.y, heat_loss: s.heat_loss, count: s.count, dir: s.dir}
}

func (s step_t) FaceRight() step_t {
	return step_t{x: s.x, y: s.y, heat_loss: s.heat_loss, count: 0, dir: turn(s.dir, right)}
}

func (s step_t) FaceLeft() step_t {
	return step_t{x: s.x, y: s.y, heat_loss: s.heat_loss, count: 0, dir: turn(s.dir, left)}
}

func (s *step_t) Forward(city_blocks *[][]city_block_t) error {

	switch s.dir {
	case right:
		s.x++
	case left:
		s.x--
	case up:
		s.y--
	case down:
		s.y++
	}

	s.count++

	if s.x < 0 || s.y < 0 || s.y >= len(*city_blocks) || s.x >= len((*city_blocks)[0]) {
		return ErrInvalidStep
	}

	s.heat_loss += (*city_blocks)[s.y][s.x].heat_loss

	return nil
}

type city_block_t struct {
	heat_loss       int
	min_heat_losses map[int]map[int]int // dir -> count -> min_dist
}

func (cb city_block_t) MinHeatLoss() int {
	first := true
	min := 0
	for _, i := range cb.min_heat_losses {
		for _, m := range i {
			if m < min || first {
				min = m
				first = false
			}
		}
	}
	return min
}

func turn(dir, rl int) int {
	switch rl {
	case right:
		switch dir {
		case up:
			return right
		case right:
			return down
		case down:
			return left
		case left:
			return up
		}
	case left:
		switch dir {
		case up:
			return left
		case right:
			return up
		case down:
			return right
		case left:
			return down
		}
	}
	log.Fatal("incorrect directions")
	return 0
}

func parseDay17Input(s string) [][]city_block_t {

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	city_blocks := make([][]city_block_t, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]city_block_t, 0)
		for _, b := range scanner.Text() {
			n, err := strconv.Atoi(string(b))
			if err != nil {
				log.Fatal(err)
			}
			city_block := city_block_t{heat_loss: n}
			city_block.min_heat_losses = make(map[int]map[int]int)
			city_block.min_heat_losses[right] = make(map[int]int)
			city_block.min_heat_losses[left] = make(map[int]int)
			city_block.min_heat_losses[up] = make(map[int]int)
			city_block.min_heat_losses[down] = make(map[int]int)
			row = append(row, city_block)
		}
		city_blocks = append(city_blocks, row)
	}

	return city_blocks
}

func dijkstraCityBlock(step step_t, city_blocks *[][]city_block_t) {

	if step.Forward(city_blocks) != nil {
		return
	}

	if v, exists := (*city_blocks)[step.y][step.x].min_heat_losses[step.dir][step.count]; exists {
		if step.heat_loss >= v {
			return
		}
	}

	for sc := step.count; sc <= max_dir; sc++ {
		if v, exists := (*city_blocks)[step.y][step.x].min_heat_losses[step.dir][sc]; exists {
			(*city_blocks)[step.y][step.x].min_heat_losses[step.dir][sc] = min(v, step.heat_loss)
		} else {
			(*city_blocks)[step.y][step.x].min_heat_losses[step.dir][sc] = step.heat_loss
		}
	}

	if step.heat_loss > best {
		return
	}

	// check if we're at the end
	if step.y == len(*city_blocks)-1 && step.x == len((*city_blocks)[0])-1 {
		if (*city_blocks)[step.y][step.x].MinHeatLoss() < best {
			best = step.heat_loss
			if show_best {
				fmt.Println(best)
			}
		}
		return
	}

	// continue forward (if applicable)
	if step.count < max_dir {
		dijkstraCityBlock(step.Duplicate(), city_blocks)
	}

	// turn (if applicable)
	if step.count >= min_dir {
		steps := []step_t{step.FaceRight(), step.FaceLeft()}
		// rand.Shuffle(len(steps), func(i, j int) { steps[i], steps[j] = steps[j], steps[i] }) // randomly turn left or right
		for _, step := range steps {
			if step.heat_loss < best {
				dijkstraCityBlock(step, city_blocks)
			}
		}
	}
}

func day17Search(min_dir0, max_dir0, threshold int, s string) int {
	city_blocks := parseDay17Input(s)
	min_dir = min_dir0
	max_dir = max_dir0
	best = threshold // this was whittled down over consecutive runs
	for _, dir := range []int{left, right, up, down} {
		dijkstraCityBlock(step_t{x: 0, y: 0, heat_loss: 0, count: 0, dir: dir}, &city_blocks)
	}
	return city_blocks[len(city_blocks)-1][len(city_blocks[0])-1].MinHeatLoss()
}

func main() {
	fmt.Println("Day 17")

	show_best = true

	input := filepath.Join("input", "input17.txt")

	part1 := day17Search(0, 3, math.MaxInt, input)
	// part1 := day17Search(0, 3, 797, input)
	fmt.Printf("part1: %v\n", part1)

	part2 := day17Search(4, 10, math.MaxInt, input)
	// part2 := day17Search(4, 10, 960, input)
	fmt.Printf("part2: %v\n", part2)
}
