package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type platform_t struct {
	pos [][]byte
}

func (p platform_t) String() string {
	output := ""
	for _, row := range p.pos {
		output += string(row) + "\n"
	}
	return output
}

var ErrNoPeriod error = errors.New("period not found")

func (p *platform_t) FindNextPeriodInCol(start_row, col int) (int, error) {

	for newrow := start_row; newrow < len(p.pos); newrow++ {
		if p.pos[newrow][col] == '.' {
			return newrow, nil
		}
	}

	return 0, ErrNoPeriod
}

func (p *platform_t) FindPrevPeriodInCol(start_row, col int) (int, error) {

	for newrow := start_row; newrow >= 0; newrow-- {
		if p.pos[newrow][col] == '.' {
			return newrow, nil
		}
	}

	return 0, ErrNoPeriod
}

func (p *platform_t) FindNextPeriodInRow(start_col, row int) (int, error) {

	for newcol := start_col; newcol < len(p.pos[0]); newcol++ {
		if p.pos[row][newcol] == '.' {
			return newcol, nil
		}
	}

	return 0, ErrNoPeriod
}

func (p *platform_t) FindPrevPeriodInRow(start_col, row int) (int, error) {

	for newcol := start_col; newcol >= 0; newcol-- {
		if p.pos[row][newcol] == '.' {
			return newcol, nil
		}
	}

	return 0, ErrNoPeriod
}

func (p *platform_t) TiltNorth() {

	var err error

	ncols := len(p.pos[0])
	nrows := len(p.pos)

	for col := 0; col < ncols; col++ {
		swap_row := 0

		for row := 0; row < nrows; row++ {
			if p.pos[row][col] == '#' {
				// find next period
				swap_row, err = p.FindNextPeriodInCol(row+1, col)
				if err != nil {
					break
				}
				row = swap_row
			} else if p.pos[row][col] == 'O' {
				p.pos[row][col] = '.'
				p.pos[swap_row][col] = 'O'
				swap_row, err = p.FindNextPeriodInCol(swap_row, col)

				if err != nil {
					break
				}
				row = swap_row
			}
		}
	}
}

func (p *platform_t) TiltWest() {

	var err error

	ncols := len(p.pos[0])
	nrows := len(p.pos)

	for row := 0; row < nrows; row++ {
		swap_col := 0

		for col := 0; col < ncols; col++ {
			if p.pos[row][col] == '#' {
				// find next period
				swap_col, err = p.FindNextPeriodInRow(col+1, row)
				if err != nil {
					break
				}
				col = swap_col
			} else if p.pos[row][col] == 'O' {
				p.pos[row][col] = '.'
				p.pos[row][swap_col] = 'O'
				swap_col, err = p.FindNextPeriodInRow(swap_col, row)

				if err != nil {
					break
				}
				col = swap_col
			}
		}
	}
}

func (p *platform_t) TiltSouth() {

	var err error

	ncols := len(p.pos[0])
	nrows := len(p.pos)

	for col := 0; col < ncols; col++ {
		swap_row := nrows - 1

		for row := nrows - 1; row >= 0; row-- {
			if p.pos[row][col] == '#' {
				// find next period
				swap_row, err = p.FindPrevPeriodInCol(row-1, col)
				if err != nil {
					break
				}
				row = swap_row
			} else if p.pos[row][col] == 'O' {
				p.pos[row][col] = '.'
				p.pos[swap_row][col] = 'O'
				swap_row, err = p.FindPrevPeriodInCol(swap_row, col)

				if err != nil {
					break
				}
				row = swap_row
			}
		}
	}
}

func (p *platform_t) TiltEast() {

	var err error

	ncols := len(p.pos[0])
	nrows := len(p.pos)

	for row := 0; row < nrows; row++ {
		swap_col := ncols - 1

		for col := swap_col; col >= 0; col-- {
			if p.pos[row][col] == '#' {
				// find next period
				swap_col, err = p.FindPrevPeriodInRow(col-1, row)
				if err != nil {
					break
				}
				col = swap_col
			} else if p.pos[row][col] == 'O' {
				p.pos[row][col] = '.'
				p.pos[row][swap_col] = 'O'
				swap_col, err = p.FindPrevPeriodInRow(swap_col, row)

				if err != nil {
					break
				}
				col = swap_col
			}
		}
	}
}

func (p *platform_t) Cycle() {
	p.TiltNorth()
	p.TiltWest()
	p.TiltSouth()
	p.TiltEast()
}

func (p *platform_t) ComputeLoad() int {
	load := 0
	nrows := len(p.pos)
	for i, row := range p.pos {
		row_sum := 0
		for _, b := range row {
			if b == 'O' {
				row_sum++
			}
		}
		load += row_sum * (nrows - i)
	}
	return load
}

func main() {
	fmt.Println("Day 14")

	p1start := time.Now()

	file, err := os.Open(filepath.Join("input", "input14.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var platform platform_t

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]byte, 0)
		for _, b := range scanner.Text() {
			row = append(row, byte(b))
		}
		platform.pos = append(platform.pos, row)
	}

	platform.TiltNorth()
	part1 := platform.ComputeLoad()
	p1end := time.Now()
	p1dur := p1end.Sub(p1start)

	fmt.Printf("part1: %v (%v)\n", part1, p1dur) // 109661

	p2start := time.Now()

	ctl := make(map[int]int)     // maps cycles remaining to load
	ltc := make(map[int]([]int)) // maps loads to remaining cycles

	var part2 int

	for i := 1000000000 - 1; i >= 0; i-- {
		platform.Cycle()
		load := platform.ComputeLoad()
		ctl[i] = load
		ltc[load] = append(ltc[load], i)

		if len(ltc[load]) > 4 {
			n := len(ltc[load])
			d := ltc[load][n-2] - ltc[load][n-1]

			if ltc[load][n-3]-ltc[load][n-1] != 2*d {
				continue
			}

			if ltc[load][n-4]-ltc[load][n-1] != 3*d {
				continue
			}

			for i := 0; i < ltc[load][n-1]%d; i++ {
				platform.Cycle()
			}

			fmt.Printf("part2: %v\n", platform.ComputeLoad())
			break
		}
	}

	p2end := time.Now()
	p2dur := p2end.Sub(p2start)

	fmt.Printf("part2: %v (%v)\n", part2, p2dur)
}
