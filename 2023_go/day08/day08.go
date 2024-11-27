package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type Node struct {
	left  string
	right string
}

var Map map[string]Node
var instructions string

func init() {
	Map = make(map[string]Node)

	file, err := os.Open(filepath.Join("input", "input8.txt"))
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions = scanner.Text()
	scanner.Scan()

	// var locs []string

	re := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)
	for scanner.Scan() {
		r := re.FindStringSubmatch(scanner.Text())
		Map[r[1]] = Node{left: r[2], right: r[3]}

		// if r[1][2] == 'A' {
		// 	locs = append(locs, r[1])
		// }
	}
}

func part1() {
	count := 0

	for loc := "AAA"; loc != "ZZZ"; count++ {
		dir := instructions[count%len(instructions)]
		if dir == 'L' {
			loc = Map[loc].left
		} else {
			loc = Map[loc].right
		}
	}

	fmt.Println("Part 1:", count)
}

func part2() {
	locs := make([]string, 0)

	for loc := range Map {
		if loc[2] == 'A' {
			locs = append(locs, loc)
		}
	}

	for i, loc := range locs {
		step := 0
		for {
			dir := instructions[step%len(instructions)]
			if dir == 'L' {
				loc = Map[loc].left
			} else {
				loc = Map[loc].right
			}
			step++

			if loc[2] == 'Z' {
				fmt.Printf("%d: %d\n", i, step)
				break
			}
		}
	}

	fmt.Println("Part 2: By happy coincidence, the answer to this is the least common multiple of the numbers printed above.")
}

func main() {
	fmt.Println("Day 8")
	part1()
	part2()
}
