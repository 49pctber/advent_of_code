package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type reindeer_t struct {
	speed    int
	fly_dur  int
	rest_dur int
}

func (r reindeer_t) Distance(race_dur int) int {
	period := r.fly_dur + r.rest_dur
	return r.speed * (race_dur/period*r.fly_dur + min(race_dur%period, r.fly_dur))
}

func main() {
	race_dur := 2503

	re := regexp.MustCompile(`^(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.$`)

	file, err := os.Open(`input\input14.txt`)
	if err != nil {
		log.Fatal(err)
	}

	part1 := 0
	var reindeer []reindeer_t
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := re.FindStringSubmatch(scanner.Text())
		var speed, fly_dur, rest_dur int
		var err error
		if speed, err = strconv.Atoi(r[2]); err != nil {
			log.Fatal(err)
		}
		if fly_dur, err = strconv.Atoi(r[3]); err != nil {
			log.Fatal(err)
		}
		if rest_dur, err = strconv.Atoi(r[4]); err != nil {
			log.Fatal(err)
		}
		reindeer = append(reindeer, reindeer_t{speed: speed, fly_dur: fly_dur, rest_dur: rest_dur})
	}

	for _, r := range reindeer {
		part1 = max(part1, r.Distance(race_dur))
	}

	fmt.Printf("part1: %v\n", part1)

	points := make([]int, len(reindeer))
	for t := 1; t <= race_dur; t++ {
		maxdist := 0
		ridx := -1
		for i := range reindeer {
			if dist := reindeer[i].Distance(t); dist > maxdist {
				ridx = i
				maxdist = dist
			}
		}
		points[ridx]++
	}

	part2 := slices.Max(points)
	fmt.Printf("part2: %v\n", part2)
}
