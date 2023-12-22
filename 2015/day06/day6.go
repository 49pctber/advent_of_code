package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	toggle = iota
	on
	off
)

func main() {
	fmt.Println("Day 6")

	file, err := os.Open(`input\input6.txt`)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`^.* (\d+),(\d+) through (\d+),(\d+)$`)

	var grid [1000][1000]bool
	var brightness [1000][1000]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		r := re.FindStringSubmatch(scanner.Text())

		x1, err := strconv.Atoi(r[1])
		if err != nil {
			log.Fatal(err)
		}
		y1, err := strconv.Atoi(r[2])
		if err != nil {
			log.Fatal(err)
		}
		x2, err := strconv.Atoi(r[3])
		if err != nil {
			log.Fatal(err)
		}
		y2, err := strconv.Atoi(r[4])
		if err != nil {
			log.Fatal(err)
		}
		var mode int
		switch {
		case scanner.Text()[:6] == "toggle":
			mode = toggle
		case scanner.Text()[:7] == "turn on":
			mode = on
		case scanner.Text()[:8] == "turn off":
			mode = off
		default:
			log.Fatal("error parsing file")
		}

		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				switch mode {
				case toggle:
					grid[i][j] = !grid[i][j]
					brightness[i][j] += 2
				case on:
					grid[i][j] = true
					brightness[i][j] += 1
				case off:
					grid[i][j] = false
					if brightness[i][j] > 0 {
						brightness[i][j] -= 1
					}
				}
			}
		}
	}

	count := 0
	total_brightness := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] {
				count++
			}
			total_brightness += brightness[i][j]
		}
	}

	fmt.Println(count)            // 543903
	fmt.Println(total_brightness) // 14687245
}
