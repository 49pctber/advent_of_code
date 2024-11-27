package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

const (
	dir_N = iota
	dir_E
	dir_W
	dir_S
)

func Distance1(input string) int {

	instructions := strings.Split(input, ", ")

	x := 0
	y := 0
	dir := dir_N

	for _, i := range instructions {
		var turn byte = i[0]
		var dist int
		fmt.Sscanf(i[1:], "%d", &dist)

		switch turn {
		case 'L':
			switch dir {
			case dir_N:
				dir = dir_W
			case dir_E:
				dir = dir_N
			case dir_W:
				dir = dir_S
			case dir_S:
				dir = dir_E
			}
		case 'R':
			switch dir {
			case dir_N:
				dir = dir_E
			case dir_E:
				dir = dir_S
			case dir_W:
				dir = dir_N
			case dir_S:
				dir = dir_W
			}
		}

		switch dir {
		case dir_N:
			y += dist
		case dir_E:
			x += dist
		case dir_W:
			x -= dist
		case dir_S:
			y -= dist
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func checkVisits(visits *map[int]map[int]interface{}, x, y int) bool {
	if _, ok := (*visits)[x]; !ok {
		(*visits)[x] = make(map[int]interface{})
	}

	if _, ok := (*visits)[x][y]; !ok {
		(*visits)[x][y] = nil
		return false
	}

	return true
}

func Distance2(input string) int {

	instructions := strings.Split(input, ", ")

	x := 0
	y := 0
	dir := dir_N

	visits := make(map[int](map[int]interface{}))
	visits[0] = make(map[int]interface{})
	visits[0][0] = nil

outer:
	for _, i := range instructions {
		var turn byte = i[0]
		var dist int
		fmt.Sscanf(i[1:], "%d", &dist)

		switch turn {
		case 'L':
			switch dir {
			case dir_N:
				dir = dir_W
			case dir_E:
				dir = dir_N
			case dir_W:
				dir = dir_S
			case dir_S:
				dir = dir_E
			}
		case 'R':
			switch dir {
			case dir_N:
				dir = dir_E
			case dir_E:
				dir = dir_S
			case dir_W:
				dir = dir_N
			case dir_S:
				dir = dir_W
			}
		}

		for i := 0; i < dist; i++ {
			switch dir {
			case dir_N:
				y += 1
			case dir_E:
				x += 1
			case dir_W:
				x -= 1
			case dir_S:
				y -= 1
			}

			visited := checkVisits(&visits, x, y)
			if visited {
				break outer
			}
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

//go:embed input/input.txt
var input string

func main() {
	fmt.Printf("Part 1: %d\n", Distance1(input))
	fmt.Printf("Part 2: %d\n", Distance2(input))
}
