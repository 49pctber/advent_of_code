package main

import (
	"fmt"
	"math"
)

func ComputeDistance(hold int, time int) int {
	return (time - hold) * hold
}

func main() {
	fmt.Println("Day 6")

	var times []int = []int{41, 96, 88, 94}
	var distances []int = []int{214, 1789, 1127, 1055}

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

	fmt.Println("Prod: ", prod)

	var T float64 = 41968894
	var D float64 = 214178911271055

	fmt.Println(int(math.Ceil(math.Sqrt(math.Pow(T, 2.0) - 4.0*D))))

	// This solution is very fragile, but it got the job done.

}
