package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type sue_t struct {
	id         int
	attributes map[string]int
}

func (sue sue_t) CheckSue1() bool {
	if val, ok := sue.attributes["children"]; ok && val != 3 {
		return false
	}
	if val, ok := sue.attributes["cats"]; ok && val != 7 {
		return false
	}
	if val, ok := sue.attributes["samoyeds"]; ok && val != 2 {
		return false
	}
	if val, ok := sue.attributes["pomeranians"]; ok && val != 3 {
		return false
	}
	if val, ok := sue.attributes["akitas"]; ok && val != 0 {
		return false
	}
	if val, ok := sue.attributes["vizslas"]; ok && val != 0 {
		return false
	}
	if val, ok := sue.attributes["goldfish"]; ok && val != 5 {
		return false
	}
	if val, ok := sue.attributes["trees"]; ok && val != 3 {
		return false
	}
	if val, ok := sue.attributes["cars"]; ok && val != 2 {
		return false
	}
	if val, ok := sue.attributes["perfumes"]; ok && val != 1 {
		return false
	}
	return true
}

func (sue sue_t) CheckSue2() bool {
	if val, ok := sue.attributes["children"]; ok && val != 3 {
		return false
	}
	if val, ok := sue.attributes["cats"]; ok && val <= 7 {
		return false
	}
	if val, ok := sue.attributes["samoyeds"]; ok && val != 2 {
		return false
	}
	if val, ok := sue.attributes["pomeranians"]; ok && val >= 3 {
		return false
	}
	if val, ok := sue.attributes["akitas"]; ok && val != 0 {
		return false
	}
	if val, ok := sue.attributes["vizslas"]; ok && val != 0 {
		return false
	}
	if val, ok := sue.attributes["goldfish"]; ok && val >= 5 {
		return false
	}
	if val, ok := sue.attributes["trees"]; ok && val <= 3 {
		return false
	}
	if val, ok := sue.attributes["cars"]; ok && val != 2 {
		return false
	}
	if val, ok := sue.attributes["perfumes"]; ok && val != 1 {
		return false
	}
	return true
}

func main() {
	re := regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)

	file, err := os.Open(`input.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sues := []sue_t{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := re.FindStringSubmatch(scanner.Text())
		sue := sue_t{attributes: make(map[string]int, 0)}
		sue.id, _ = strconv.Atoi(r[1])
		sue.attributes[r[2]], _ = strconv.Atoi(r[3])
		sue.attributes[r[4]], _ = strconv.Atoi(r[5])
		sue.attributes[r[6]], _ = strconv.Atoi(r[7])
		sues = append(sues, sue)
	}

	for _, sue := range sues {
		if sue.CheckSue1() {
			fmt.Println("Part 1:", sue.id)
		}

		if sue.CheckSue2() {
			fmt.Println("Part 2:", sue.id)
		}
	}
}
