package main

import (
	"fmt"
	"os"
	"strings"

	_ "embed"
)

type Screen struct {
	nrows  int
	ncols  int
	pixels [][]bool
}

func NewScreen(nrows, ncols int) *Screen {
	s := &Screen{nrows: nrows, ncols: ncols}
	s.pixels = make([][]bool, nrows)
	for i := range s.pixels {
		s.pixels[i] = make([]bool, ncols)
	}
	return s
}

func (s Screen) String() string {

	ss := make([]string, 0)
	for _, row := range s.pixels {
		srow := ""
		for _, col := range row {
			if col {
				srow += "#"
			} else {
				srow += "."
			}
		}
		ss = append(ss, srow)
	}

	return strings.Join(ss, "\n")
}

func (s Screen) CountLit() int {
	count := 0
	for _, row := range s.pixels {
		for _, col := range row {
			if col {
				count += 1
			}
		}
	}
	return count
}

func (s *Screen) Rect(nrows, ncols int) {
	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			s.pixels[i][j] = true
		}
	}
}

func (s *Screen) RotateRow(row, shift int) {

	// compute new row
	newrow := make([]bool, s.ncols)
	for i := 0; i < s.ncols; i++ {
		j := (i + shift) % s.ncols
		newrow[j] = s.pixels[row][i]
	}

	// copy newrow to actual row
	for i := 0; i < s.ncols; i++ {
		s.pixels[row][i] = newrow[i]
	}
}

func (s *Screen) RotateCol(col, shift int) {
	// compute new column
	newcol := make([]bool, s.nrows)
	for i := 0; i < s.nrows; i++ {
		j := (i + shift) % s.nrows
		newcol[j] = s.pixels[i][col]
	}

	// copy newrow to actual row
	for i := 0; i < s.nrows; i++ {
		s.pixels[i][col] = newcol[i]
	}
}

func LoadInput(fname string) ([]string, error) {
	b, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(b), "\r\n"), nil
}

func ParseInstructions(nrows, ncols int, instructions []string) (*Screen, error) {
	s := NewScreen(nrows, ncols)

	var a, b int

	for _, instr := range instructions {
		instr = strings.TrimSpace(instr)
		if n, err := fmt.Sscanf(instr, "rect %dx%d", &a, &b); n == 2 && err == nil {
			s.Rect(b, a)
		} else if n, err := fmt.Sscanf(instr, "rotate column x=%d by %d", &a, &b); n == 2 && err == nil {
			s.RotateCol(a, b)
		} else if n, err := fmt.Sscanf(instr, "rotate row y=%d by %d", &a, &b); n == 2 && err == nil {
			s.RotateRow(a, b)
		} else {
			continue
		}
	}

	return s, nil
}

func main() {
	instructions, err := LoadInput("input")
	if err != nil {
		panic(err)
	}

	s, err := ParseInstructions(6, 50, instructions)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %v\n", s.CountLit())
	fmt.Println(s)
}
