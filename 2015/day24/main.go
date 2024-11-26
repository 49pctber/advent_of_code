package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func nextCombination(current []int, n int) bool {
	k := len(current)

	// Find the rightmost element that can be incremented
	for i := k - 1; i >= 0; i-- {
		if current[i] < n-k+i {
			current[i]++
			// Adjust the subsequent elements
			for j := i + 1; j < k; j++ {
				current[j] = current[j-1] + 1
			}
			return true
		}
	}

	return false // No more combinations
}

func check(idxs []int, weights []int, target int) int {

	sum := 0
	for _, idx := range idxs {
		sum += weights[idx]
	}

	if sum != target {
		return math.MaxInt
	}

	// TODO check that it is possible to evenly divide remaining gifts into equal weight

	qe := 1
	for _, idx := range idxs {
		qe *= weights[idx]
	}

	return qe
}

func main() {

	fmt.Println("Day 24")

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var weights []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		weight, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		weights = append(weights, weight)
	}

	sort.Slice(weights, func(i, j int) bool {
		return weights[i] < weights[j]
	})

	cumulative := 0
	for _, weight := range weights {
		cumulative += weight
	}

	if cumulative%3 != 0 {
		panic("cumulative weight is not divisible by three")
	}

	if cumulative%4 != 0 {
		panic("cumulative weight is not divisible by four")
	}

	// part 1
	minqe := math.MaxInt
	target := cumulative / 3

	for i := 1; i < len(weights); i++ {
		idxs := make([]int, i)
		for n := range idxs {
			idxs[n] = n
		}

		for nextCombination(idxs, len(weights)) {
			minqe = min(minqe, check(idxs, weights, target))
		}

		if minqe != math.MaxInt {
			break
		}
	}

	fmt.Println("Part 1:", minqe)

	// part 2
	minqe = math.MaxInt
	target = cumulative / 4

	for i := 1; i < len(weights); i++ {
		idxs := make([]int, i)
		for n := range idxs {
			idxs[n] = n
		}

		for nextCombination(idxs, len(weights)) {
			minqe = min(minqe, check(idxs, weights, target))
		}

		if minqe != math.MaxInt {
			break
		}
	}

	fmt.Println("Part 2:", minqe)
}
