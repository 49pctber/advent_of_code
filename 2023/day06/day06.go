package main

import (
	"fmt"
	"math"
	"strconv"
)

func ComputeDistance(hold int, time int) int {
	return (time - hold) * hold
}

func main() {
	fmt.Println("Day 6")

	var times []int = []int{55, 99, 97, 93}
	var distances []int = []int{401, 1485, 2274, 1405}

	// for _, i := range []int{0, 1, 2, 3, 4, 5, 6, 7} {
	// 	fmt.Println(i, ComputeDistance(i, 7))
	// }

	prod := 1

	for i := range times {
		n := 0
		for j := 0; j < times[i]; j++ {
			if ComputeDistance(j, times[i]) > distances[i] {
				n++
			}
		}
		prod *= n
	}

	fmt.Println("Part 1:", prod)

	Tstr := ""
	for _, n := range times {
		Tstr += strconv.Itoa(n)
	}

	Dstr := ""
	for _, n := range distances {
		Dstr += strconv.Itoa(n)
	}

	T, _ := strconv.ParseFloat(Tstr, 64)
	D, _ := strconv.ParseFloat(Dstr, 64)

	fmt.Println("Part 2:", int(math.Ceil(math.Sqrt(math.Pow(T, 2.0)-4.0*D))))

}
