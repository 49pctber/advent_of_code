package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Recurse(containers []int, remaining int, idx int) int {

	// base case
	if remaining == 0 {
		return 1
	} else if idx == len(containers) {
		return 0
	}

	// inductive case
	sum := 0
	if containers[idx] <= remaining {
		sum += Recurse(containers, remaining-containers[idx], idx+1)
	}
	sum += Recurse(containers, remaining, idx+1)
	return sum
}

func Start(containers []int, target int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(containers)))

	// fmt.Printf("containers: %v\n", containers)

	return Recurse(containers, target, 0)
}

func ParseInput(s string) []int {
	file, err := os.Open(`input.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var containers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		containers = append(containers, n)
	}

	return containers
}

var results map[int]int = make(map[int]int)

func Recurse2(containers []int, remaining int, count int, idx int) {

	// base case
	if remaining == 0 {
		results[count]++
		return
	} else if idx == len(containers) {
		return
	}

	// inductive case
	if containers[idx] <= remaining {
		Recurse2(containers, remaining-containers[idx], count+1, idx+1)
	}
	Recurse2(containers, remaining, count, idx+1)
}

func Start2(containers []int, target int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(containers)))

	// fmt.Printf("containers: %v\n", containers)

	Recurse2(containers, target, 0, 0)

	keys := make([]int, 0, len(results))
	for k := range results {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	return results[keys[0]]
}

func main() {
	containers := ParseInput(`input.txt`)
	target := 150
	part1 := Start(containers, target)
	fmt.Printf("part1: %v\n", part1)

	part2 := Start2(containers, target)
	fmt.Printf("part2: %v\n", part2)
}
