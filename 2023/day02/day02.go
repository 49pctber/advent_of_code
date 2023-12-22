package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Subset struct {
	red   int
	green int
	blue  int
}

func (s *Subset) Populate(subset_str string) {
	re_red := regexp.MustCompile(`(\d+) red`)
	re_blue := regexp.MustCompile(`(\d+) blue`)
	re_green := regexp.MustCompile(`(\d+) green`)

	r := re_red.FindStringSubmatch(subset_str)
	if r != nil {
		s.red, _ = strconv.Atoi(r[1])
	}

	r = re_blue.FindStringSubmatch(subset_str)
	if r != nil {
		s.blue, _ = strconv.Atoi(r[1])
	}

	r = re_green.FindStringSubmatch(subset_str)
	if r != nil {
		s.green, _ = strconv.Atoi(r[1])
	}
}

func (s *Subset) Valid() bool {
	return s.red <= 12 && s.blue <= 14 && s.green <= 13
}

func (s *Subset) Print() {
	fmt.Printf("R: %d, B: %d, G: %d (%t)\n", s.red, s.blue, s.green, s.Valid())
}

type Game struct {
	id      int
	subsets []Subset
}

func (game *Game) Print() {
	fmt.Printf("Game #%d (%t) \n", game.id, game.Valid())
	for _, subset := range game.subsets {
		subset.Print()
	}
}

func (game *Game) Valid() bool {
	valid := true
	for _, subset := range game.subsets {
		valid = valid && subset.Valid()
	}
	return valid
}

func (game *Game) Power() int {
	min_red := 0
	min_blue := 0
	min_green := 0

	for _, subset := range game.subsets {
		if subset.red > min_red {
			min_red = subset.red
		}

		if subset.blue > min_blue {
			min_blue = subset.blue
		}

		if subset.green > min_green {
			min_green = subset.green
		}
	}

	// compute product of minima
	return min_red * min_blue * min_green
}

func main() {
	fmt.Println("Day 2")
	input, err := os.ReadFile("input/input2.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(input), "\n")

	re := regexp.MustCompile(`^Game (\d+): (.*)$`)

	n_valid := 0
	sum := 0
	sum_power := 0

	for _, row := range rows {
		// fmt.Println(i, row)
		matches := re.FindAllStringSubmatch(row, -1)

		for _, match := range matches {
			id, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			game := Game{id: id}
			subsets := strings.Split(match[2], ";")

			for _, subset_str := range subsets {
				var subset Subset
				subset.Populate(subset_str)
				game.subsets = append(game.subsets, subset)
			}

			if game.Valid() {
				n_valid++
				sum += game.id
			}

			sum_power += game.Power()

			game.Print()
		}
	}

	fmt.Println("Number of valid games:", n_valid)
	fmt.Println("Sum of indices:", sum)
	fmt.Println("Sum of powers:", sum_power)
}
