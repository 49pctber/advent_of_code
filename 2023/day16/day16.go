package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	up int = iota
	down
	left
	right
)

type beam_t struct {
	row int
	col int
	dir int
}

func (b beam_t) String() string {
	lut := map[int]string{up: "^", down: "v", right: ">", left: "<"}
	return fmt.Sprintf("(%d,%d)%s", b.row, b.col, lut[b.dir])
}

type wall_t struct {
	location  [][]byte
	energized [][]bool
	beams     []beam_t
	memo      map[string]struct{}
}

func (w wall_t) CountEnergizedTiles() int {
	count := 0
	for i := range w.energized {
		for j := range w.energized[i] {
			if w.energized[i][j] {
				count++
				w.energized[i][j] = false
			}
		}
	}
	return count
}

func (w *wall_t) propagateBeam() {
	var beam *beam_t = &(w.beams[0]) // TODO may be a bug here
	w.beams = w.beams[1:]

	if _, ok := w.memo[beam.String()]; ok {
		return
	} else {
		w.memo[beam.String()] = struct{}{}
	}

	for {
		// check for out of bounds
		if beam.row < 0 || beam.row >= len(w.location) {
			return
		}

		if beam.col < 0 || beam.col >= len(w.location[beam.row]) {
			return
		}

		// mark energy
		w.energized[beam.row][beam.col] = true

		// check direction
		switch w.location[beam.row][beam.col] {
		case '\\':
			switch beam.dir {
			case right:
				beam.dir = down
			case left:
				beam.dir = up
			case down:
				beam.dir = right
			case up:
				beam.dir = left
			}
		case '/':
			switch beam.dir {
			case right:
				beam.dir = up
			case left:
				beam.dir = down
			case down:
				beam.dir = left
			case up:
				beam.dir = right
			}
		case '-':
			switch beam.dir {
			case down:
				w.beams = append(w.beams, beam_t{row: beam.row, col: beam.col + 1, dir: right})
				w.beams = append(w.beams, beam_t{row: beam.row, col: beam.col - 1, dir: left})
				return
			case up:
				w.beams = append(w.beams, beam_t{row: beam.row, col: beam.col + 1, dir: right})
				w.beams = append(w.beams, beam_t{row: beam.row, col: beam.col - 1, dir: left})
				return
			}
		case '|':
			switch beam.dir {
			case right:
				w.beams = append(w.beams, beam_t{row: beam.row - 1, col: beam.col, dir: up})
				w.beams = append(w.beams, beam_t{row: beam.row + 1, col: beam.col, dir: down})
				return
			case left:
				w.beams = append(w.beams, beam_t{row: beam.row - 1, col: beam.col, dir: up})
				w.beams = append(w.beams, beam_t{row: beam.row + 1, col: beam.col, dir: down})
				return
			}
		case '.':

		default:
			log.Fatal("THERE")
		}

		// move beam
		switch beam.dir {
		case up:
			beam.row--
		case down:
			beam.row++
		case right:
			beam.col++
		case left:
			beam.col--
		default:
			log.Fatal("HERE")
		}
	}
}

func (w *wall_t) PropagateBeams() {
	w.memo = make(map[string]struct{})
	for i := 0; len(w.beams) > 0; i++ {
		w.propagateBeam()
	}
}

func main() {
	file, err := os.Open(filepath.Join("input", "input16.txt"))
	if err != nil {
		log.Fatal(err)
	}

	wall := wall_t{beams: []beam_t{{row: 0, col: 0, dir: right}}}

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		wall.location = append(wall.location, []byte{})
		wall.energized = append(wall.energized, []bool{})
		for _, b := range scanner.Text() {
			wall.location[row] = append(wall.location[row], byte(b))
			wall.energized[row] = append(wall.energized[row], false)
		}
		row++
	}
	// fmt.Printf("wall.location: %v\n", wall.location)

	wall.PropagateBeams()
	part1 := wall.CountEnergizedTiles()
	fmt.Printf("part1: %v\n", part1) // 7562

	part2 := 0
	for i := range wall.location {
		wall.beams = append(wall.beams, beam_t{row: i, col: 0, dir: right})
		wall.PropagateBeams()
		part2 = max(part2, wall.CountEnergizedTiles())

		wall.beams = append(wall.beams, beam_t{row: i, col: len(wall.location[0]) - 1, dir: left})
		wall.PropagateBeams()
		part2 = max(part2, wall.CountEnergizedTiles())
	}
	for j := range wall.location[0] {
		wall.beams = append(wall.beams, beam_t{row: 0, col: j, dir: down})
		wall.PropagateBeams()
		part2 = max(part2, wall.CountEnergizedTiles())

		wall.beams = append(wall.beams, beam_t{row: len(wall.location) - 1, col: j, dir: up})
		wall.PropagateBeams()
		part2 = max(part2, wall.CountEnergizedTiles())
	}

	fmt.Printf("part2: %v\n", part2) // 7791 < ans
}
