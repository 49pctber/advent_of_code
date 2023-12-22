package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Galaxy struct {
	x int
	y int
}

func (g *Galaxy) String() string {
	return fmt.Sprintf("(%d,%d)", g.x, g.y)
}

func galaxyDistance(g1, g2 Galaxy, f int) int {
	dist := 0

	xmin := min(g1.x, g2.x)
	xmax := max(g1.x, g2.x)
	ymin := min(g1.y, g2.y)
	ymax := max(g1.y, g2.y)

	for x := xmin; x < xmax; x++ {
		if double_columns[x] {
			dist += f
		} else {
			dist += 1
		}
	}

	for y := ymin; y < ymax; y++ {
		if double_rows[y] {
			dist += f
		} else {
			dist += 1
		}
	}
	return dist
}

var double_columns map[int]bool
var double_rows map[int]bool

func init() {
	double_columns = make(map[int]bool)
	double_rows = make(map[int]bool)
}

func main() {
	fmt.Println("Day 11")

	file, err := os.Open(`input\input11.txt`)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	i := 0

	var galaxies []Galaxy

	for scanner.Scan() {
		row := scanner.Text()
		double_rows[i] = true
		for j, c := range row {
			double_columns[j] = true
			if c == '#' {
				galaxies = append(galaxies, Galaxy{x: j, y: i})
			}
		}
		i++
	}

	for _, g := range galaxies {
		double_rows[g.y] = false
		double_columns[g.x] = false
	}

	dist_sum := 0
	dist_sum2 := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			dist_sum += galaxyDistance(galaxies[i], galaxies[j], 2)
			dist_sum2 += galaxyDistance(galaxies[i], galaxies[j], 1000000)
		}
	}

	fmt.Printf("dist_sum: %v\n", dist_sum)   // 9370588
	fmt.Printf("dist_sum2: %v\n", dist_sum2) // 746207878188
}
