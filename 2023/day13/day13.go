package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type valley struct {
	rows []int
	cols []int
}

func (v valley) String() string {
	return fmt.Sprintf("Rows: %v, Cols: %v", v.rows, v.cols)
}

func (v *valley) ProcessString(s string) {

	// update columns
	for i, c := range s {
		if c == '#' {
			v.cols[i] += 1 << len(v.rows)
		}
	}

	// update rows
	val := 0
	for i, c := range s {
		if c == '#' {
			val += 1 << i
		}
	}
	v.rows = append(v.rows, val)

}

func (v *valley) ComputeValue() (int, error) {

	for i := 1; i < len(v.rows); i++ {
		if v.rows[i-1] == v.rows[i] {
			valid := true
			for j := 1; i-1-j >= 0 && i+j < len(v.rows); j++ {
				if v.rows[i-1-j] != v.rows[i+j] {
					valid = false
					break
				}
			}
			if valid {
				return 100 * i, nil
			}
		}
	}

	for i := 1; i < len(v.cols); i++ {
		if v.cols[i-1] == v.cols[i] {
			valid := true
			for j := 1; i-1-j >= 0 && i+j < len(v.cols); j++ {
				if v.cols[i-1-j] != v.cols[i+j] {
					valid = false
					break
				}
			}
			if valid {
				return i, nil
			}
		}
	}

	return 0, fmt.Errorf("no reflection found: rows: %v, cols: %v", v.rows, v.cols)

}

func (v *valley) ComputeNewValue(og_val int) (int, error) {

	for i := 1; i < len(v.rows); i++ {
		if v.rows[i-1] == v.rows[i] {
			valid := true
			for j := 1; i-1-j >= 0 && i+j < len(v.rows); j++ {
				if v.rows[i-1-j] != v.rows[i+j] {
					valid = false
					break
				}
			}
			new_val := 100 * i
			if valid && new_val != og_val {
				return new_val, nil
			}
		}
	}

	for i := 1; i < len(v.cols); i++ {
		if v.cols[i-1] == v.cols[i] {
			valid := true
			for j := 1; i-1-j >= 0 && i+j < len(v.cols); j++ {
				if v.cols[i-1-j] != v.cols[i+j] {
					valid = false
					break
				}
			}
			new_val := i
			if valid && new_val != og_val {
				return new_val, nil
			}
		}
	}

	return 0, fmt.Errorf("no reflection found: rows: %v, cols: %v", v.rows, v.cols)

}

func (v *valley) Smudge() (int, error) {
	var og_val int
	var err error

	if og_val, err = v.ComputeValue(); err != nil {
		return 0, err
	}

	for i := 0; i < len(v.rows); i++ {
		for j := 0; j < len(v.cols); j++ {

			v.rows[i] ^= 1 << j
			v.cols[j] ^= 1 << i
			val, err := v.ComputeNewValue(og_val)
			if err == nil && val != og_val {
				return val, nil
			}
			v.rows[i] ^= 1 << j
			v.cols[j] ^= 1 << i
		}
	}

	return 0, fmt.Errorf("no new values found")
}

func main() {
	fmt.Println("Day 13")

	file, err := os.Open(`input\input13.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read in grids
	scanner := bufio.NewScanner(file)
	new_valley := true
	var valleys []valley
	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			new_valley = true
			continue
		}

		if new_valley {
			new_valley = false
			valley := valley{}
			valley.cols = make([]int, len(input))
			valley.rows = make([]int, 0)
			valley.ProcessString(input)
			valleys = append(valleys, valley)
		} else {
			valleys[len(valleys)-1].ProcessString(input)
		}
	}

	part1 := 0
	for _, valley := range valleys {
		val, err := valley.ComputeValue()
		if err != nil {
			log.Fatal(err)
		}
		part1 += val
	}

	fmt.Printf("part1: %v\n", part1) // 37113

	part2 := 0
	for _, valley := range valleys {

		val, err := valley.Smudge()
		if err != nil {
			log.Fatal(err)
		}

		part2 += val
	}

	fmt.Printf("part2: %v\n", part2) // 30449

}
