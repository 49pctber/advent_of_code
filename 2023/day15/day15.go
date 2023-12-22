package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type label_t struct {
	label        string
	focal_length int
}

func Hash(s string) int {
	current := 0
	for _, b := range s {
		current = 17 * (current + int(b)) % 256
	}
	return current
}

// func remove(slice []int, s int) []int {
//     return append(slice[:s], slice[s+1:]...)
// }

var ErrNotFound error = errors.New("label not found")

func findLabel(s string, labels []label_t) (int, error) {
	for i, label := range labels {
		if label.label == s {
			return i, nil
		}
	}
	return 0, ErrNotFound
}

func main() {
	fmt.Println("Day 15")

	file, err := os.Open(`input\input15.txt`)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	steps := strings.Split(scanner.Text(), ",")

	part1 := 0
	for _, step := range steps {
		part1 += Hash(step)
	}
	fmt.Printf("part1: %v\n", part1) // 501680

	hashmap := make([][]label_t, 256)

	re := regexp.MustCompile(`^(\w+)([=-])(\d+)?$`)

	for _, step := range steps {
		r := re.FindStringSubmatch(step)
		h := Hash(r[1])

		if r[2] == "-" {

			i, err := findLabel(r[1], hashmap[h])
			if err == nil {
				hashmap[h] = append(hashmap[h][:i], hashmap[h][i+1:]...)
			}
		} else {
			n, err := strconv.Atoi(r[3])
			if err != nil {
				log.Fatal(err)
			}

			i, err := findLabel(r[1], hashmap[h])
			if err != nil {
				hashmap[h] = append(hashmap[h], label_t{focal_length: n, label: r[1]})
			} else {
				hashmap[h][i].focal_length = n
			}
		}
	}

	fmt.Printf("hashmap: %v\n", hashmap)

	part2 := 0
	for i, labels := range hashmap {
		for j, label := range labels {
			part2 += (i + 1) * (j + 1) * label.focal_length
		}
	}
	fmt.Printf("part2: %v\n", part2) // 74301 < ans
}
