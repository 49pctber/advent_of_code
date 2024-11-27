package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const pph1, pph2 int = 10, 11

var primes []int

func init() {
	primes = make([]int, 0)
}

func housePresents(n int) (int, int) {

	if n == 1 {
		return pph1, pph2
	}

	// factor n
	pmax := int(math.Sqrt(float64(n)))
	factors := make([]int, 0)
	r := n
	for _, p := range primes {
		if p > pmax {
			break
		}
		for r%p == 0 {
			r /= p
			factors = append(factors, p)
		}
		if r == 1 {
			break
		}
	}

	// if prime, add to list of known primes and return early
	if len(factors) == 0 {
		primes = append(primes, n)
		return pph1 * (1 + n), pph2 * (1 + n)
	}

	all := map[int]any{1: nil}
	for _, p := range factors {
		keys := make([]int, 0, len(all))
		for k := range all {
			keys = append(keys, k)
		}
		for _, q := range keys {
			all[q*p] = nil
		}
	}

	count1, count2 := 0, 0
	for x := range all {
		count1 += x
		if x*50 >= n {
			count2 += x
		}
	}

	return pph1 * count1, pph2 * count2

}

func main() {
	fmt.Println("Day 20")

	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input, err := strconv.Atoi(string(b[:len(b)-1]))
	if err != nil {
		panic(err)
	}

	part1done := false
	part2done := false

	for n := 1; true; n++ {
		npresents1, npresents2 := housePresents(n)
		if npresents1 >= input && !part1done {
			fmt.Println("Part 1:", n)
			part1done = true
			if part1done && part2done {
				break
			}
		}

		if npresents2 >= input && !part2done {
			fmt.Println("Part 2:", n)
			part2done = true
			if part1done && part2done {
				break
			}
		}

	}
}
