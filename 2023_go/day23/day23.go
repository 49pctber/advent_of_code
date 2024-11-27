package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

const (
	path int = iota
	forest
	upslope
	rightslope
	downslope
	leftslope
	up
	right
	down
	left
)

func (maze *maze_t) Up() {
	maze.row--
}

func (maze *maze_t) Right() {
	maze.col++
}

func (maze *maze_t) Down() {
	maze.row++
}

func (maze *maze_t) Left() {
	maze.col--
}

func (maze maze_t) TileHere() rune {
	return maze.tiles[maze.row][maze.col]
}

func (maze maze_t) TileUp() rune {
	if maze.row-1 < 0 {
		return '#'
	}
	return maze.tiles[maze.row-1][maze.col]
}

func (maze maze_t) TileRight() rune {
	if maze.col+1 >= len(maze.tiles[0]) {
		return '#'
	}
	return maze.tiles[maze.row][maze.col+1]
}

func (maze maze_t) TileDown() rune {
	if maze.row+1 >= len(maze.tiles) {
		return '#'
	}
	return maze.tiles[maze.row+1][maze.col]
}

func (maze maze_t) TileLeft() rune {
	if maze.col-1 < 0 {
		return '#'
	}
	return maze.tiles[maze.row][maze.col-1]
}

func (maze maze_t) TileUpVisited() bool {
	return maze.visited[maze.row-1][maze.col]
}

func (maze maze_t) TileRightVisited() bool {
	return maze.visited[maze.row][maze.col+1]
}

func (maze maze_t) TileDownVisited() bool {
	return maze.visited[maze.row+1][maze.col]
}

func (maze maze_t) TileLeftVisited() bool {
	return maze.visited[maze.row][maze.col-1]
}

type maze_t struct {
	tiles      [][]rune
	visited    [][]bool
	forks      map[int]map[int]int
	row        int
	col        int
	start_hash int
	end_hash   int
}

func (maze maze_t) String() string {
	rows := []string{}
	for _, row := range maze.tiles {
		rows = append(rows, string(row))
	}
	return strings.Join(rows, "\n")
}

func (maze maze_t) IsFork(row, col int) bool {
	if row == 0 || row == len(maze.tiles)-1 || col == 0 || col == len(maze.tiles[0])-1 {
		return false
	}

	if maze.tiles[row][col] == '#' {
		return false
	}

	count := 0

	if maze.tiles[row-1][col] != '#' {
		count++
	}
	if maze.tiles[row+1][col] != '#' {
		count++
	}
	if maze.tiles[row][col-1] != '#' {
		count++
	}
	if maze.tiles[row][col+1] != '#' {
		count++
	}

	return count >= 3
}

func (maze maze_t) AtFork() bool {
	return maze.IsFork(maze.row, maze.col)
}

func (maze maze_t) CountForks() (count int) {
	for i := range maze.tiles {
		for j := range maze.tiles[i] {
			if maze.IsFork(i, j) {
				count++
			}
		}
	}
	return count
}

func (maze maze_t) AtEnd() bool {
	return maze.row == len(maze.tiles)-1
}

func parseInput(s string, part2 bool) maze_t {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maze := maze_t{
		tiles:   make([][]rune, 0),
		visited: make([][]bool, 0),
		forks:   make(map[int]map[int]int),
	}

	scanner := bufio.NewScanner(file)
	for row := 0; scanner.Scan(); row++ {
		maze.tiles = append(maze.tiles, make([]rune, len(scanner.Text())))
		maze.visited = append(maze.visited, make([]bool, len(scanner.Text())))
		for col, r := range scanner.Text() {
			maze.tiles[row][col] = r
		}
	}

	// find start
	for i, c := range maze.tiles[0] {
		if c == '.' {
			maze.start_hash = maze.HashGivenLocation(0, i)
			maze.row = 0
			maze.col = i
			break
		}
	}

	// find end
	for i, c := range maze.tiles[len(maze.tiles)-1] {
		if c == '.' {
			maze.end_hash = maze.HashGivenLocation(len(maze.tiles)-1, i)
			break
		}
	}

	// find forks
	maze.FindForks(part2)

	return maze
}

func (maze maze_t) HashGivenLocation(row, col int) int {
	return row*len(maze.tiles[0]) + col
}

func (maze maze_t) HashLocation() int {
	return maze.HashGivenLocation(maze.row, maze.col)
}

func backwards(dir int) int {
	switch dir {
	case up:
		return down
	case down:
		return up
	case right:
		return left
	case left:
		return right
	default:
		log.Fatal("invalid direction")
		return 0
	}
}

func (maze *maze_t) HikeToFork(dist int, dir int, forward bool, backward bool, part2 bool) (int, bool, bool) {

	maze.visited[maze.row][maze.col] = true

	// base case
	if dist > 0 && (maze.AtFork() || maze.AtEnd()) {
		return dist, forward, backward
	}

	// keep walking
	var possible_dirs []int
	if part2 {
		possible_dirs = []int{up, down, left, right}
	} else {
		switch maze.TileHere() {

		case '^':
			possible_dirs = []int{up}
			switch dir {
			case down:
				forward = false
			default:
				backward = false
			}

		case '>':
			possible_dirs = []int{right}
			switch dir {
			case left:
				forward = false
			default:
				backward = false
			}

		case 'v':
			possible_dirs = []int{down}
			switch dir {
			case up:
				forward = false
			default:
				backward = false
			}

		case '<':
			possible_dirs = []int{left}
			switch dir {
			case right:
				forward = false
			default:
				backward = false
			}

		default:
			possible_dirs = []int{up, down, left, right}
		}
	}

	if dist == 0 {
		possible_dirs = []int{dir}
	}

	for _, nextdir := range possible_dirs {
		if nextdir == backwards(dir) {
			continue
		} else {
			switch nextdir {
			case down:
				if tile := maze.TileDown(); tile != '#' {
					maze.Down()
					return maze.HikeToFork(dist+1, down, forward, backward, part2)
				}
			case right:
				if tile := maze.TileRight(); tile != '#' {
					maze.Right()
					return maze.HikeToFork(dist+1, right, forward, backward, part2)
				}
			case up:
				if tile := maze.TileUp(); tile != '#' {
					maze.Up()
					return maze.HikeToFork(dist+1, up, forward, backward, part2)
				}
			case left:
				if tile := maze.TileLeft(); tile != '#' {
					maze.Left()
					return maze.HikeToFork(dist+1, left, forward, backward, part2)
				}
			default:
				log.Fatal("invalid direction")
			}
		}
	}

	// can't go this way
	// fmt.Println("can't go here", maze.row, maze.col)
	return 0, false, false
}

type search_t struct {
	row int
	col int
	dir int
}

func (maze *maze_t) FindForks(part2 bool) {

	queue := []search_t{{row: maze.row, col: maze.col, dir: down}}

	for len(queue) > 0 {
		// fmt.Println(queue)

		var next_search search_t
		next_search, queue = queue[0], queue[1:]

		// set location
		maze.row = next_search.row
		maze.col = next_search.col
		start_hash := maze.HashLocation()

		// search in direction
		dist, forward, backward := maze.HikeToFork(0, next_search.dir, true, true, part2)
		end_hash := maze.HashLocation()

		if !forward && !backward {
			continue
		}

		// add path with distance to map, with appropriate forward/backward
		if forward {
			if _, ok := maze.forks[start_hash]; !ok {
				maze.forks[start_hash] = make(map[int]int)
			}
			maze.forks[start_hash][end_hash] = dist
		}

		if maze.AtEnd() {
			continue
		}

		if backward {
			if _, ok := maze.forks[end_hash]; !ok {
				maze.forks[end_hash] = make(map[int]int)
			}
			maze.forks[end_hash][start_hash] = dist
		}

		// search for available directions, ingorning paths that have already been visited
		if maze.TileUp() != '#' && !maze.TileUpVisited() {
			queue = append(queue, search_t{row: maze.row, col: maze.col, dir: up})
		}

		if maze.TileDown() != '#' && !maze.TileDownVisited() {
			queue = append(queue, search_t{row: maze.row, col: maze.col, dir: down})
		}

		if maze.TileRight() != '#' && !maze.TileRightVisited() {
			queue = append(queue, search_t{row: maze.row, col: maze.col, dir: right})
		}

		if maze.TileLeft() != '#' && !maze.TileLeftVisited() {
			queue = append(queue, search_t{row: maze.row, col: maze.col, dir: left})
		}
	}
}

func (maze maze_t) pathLength(path []int) int {
	distance := 0
	for i := range path {
		if i == 0 {
			continue
		}
		distance += maze.forks[path[i-1]][path[i]]
	}
	return distance
}

func (maze maze_t) recurseLongestHike(path []int) int {

	next_hash := path[len(path)-1]
	if next_hash == maze.end_hash {
		length := maze.pathLength(path)
		// fmt.Println(length, path)
		return length
	}

	// add a new hash that isn't already in the path
	longest_path_length := 0
	for key := range maze.forks[next_hash] {
		// check for key already in path
		if slices.Contains(path, key) {
			continue
		}
		new_path := make([]int, len(path))
		copy(new_path, path)
		new_path = append(new_path, key)
		longest_path_length = max(longest_path_length, maze.recurseLongestHike(new_path))
	}
	return longest_path_length
}

func (maze maze_t) longestHike() int {
	path := []int{maze.start_hash}

	return maze.recurseLongestHike(path)
}

func main() {
	maze := parseInput("input23.txt", false)
	part1 := maze.longestHike()
	fmt.Printf("part1: %v\n", part1)

	maze = parseInput("input23.txt", true)
	part2 := maze.longestHike()
	fmt.Printf("part2: %v\n", part2)
}
