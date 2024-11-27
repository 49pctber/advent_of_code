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

var dists map[string]map[string]int

func init() {
	dists = make(map[string]map[string]int)
}

func ensureCityExists(city string) {
	_, ok := dists[city]
	if !ok {
		dists[city] = make(map[string]int)
	}
}

type path_t struct {
	path []string
	len  int
}

func computeMinPathLength(p path_t) int {
	if len(p.path) == 8 {
		return p.len
	}

	minlen := int(^uint(0) >> 1)
	for city := range dists {
		if !slices.Contains(p.path, city) {
			var new_path path_t
			new_path.path = make([]string, len(p.path))
			copy(new_path.path, p.path)
			new_path.path = append(new_path.path, city)
			if len(p.path) > 0 {
				new_path.len = p.len + dists[p.path[len(p.path)-1]][city]
			} else {
				new_path.len = 0
			}
			minlen = min(minlen, computeMinPathLength(new_path))
		}
	}

	return minlen
}

func computeMaxPathLength(p path_t) int {
	if len(p.path) == 8 {
		return p.len
	}

	maxlen := 0
	for city := range dists {
		if !slices.Contains(p.path, city) {
			var new_path path_t
			new_path.path = make([]string, len(p.path))
			copy(new_path.path, p.path)
			new_path.path = append(new_path.path, city)
			if len(p.path) > 0 {
				new_path.len = p.len + dists[p.path[len(p.path)-1]][city]
			} else {
				new_path.len = 0
			}
			maxlen = max(maxlen, computeMaxPathLength(new_path))
		}
	}

	return maxlen
}

func main() {
	fmt.Println("Day 9")
	file, err := os.Open(`input.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`(\w+) to (\w+) = (\d+)`)
	for scanner.Scan() {
		r := re.FindStringSubmatch(scanner.Text())
		ensureCityExists(r[1])
		ensureCityExists(r[2])
		n, err := strconv.Atoi(r[3])
		if err != nil {
			log.Fatal(err)
		}
		dists[r[1]][r[2]] = n
		dists[r[2]][r[1]] = n
	}

	minlen := computeMinPathLength(path_t{})
	fmt.Printf("minlen: %v\n", minlen) // 117

	maxlen := computeMaxPathLength(path_t{})
	fmt.Printf("minlen: %v\n", maxlen) // 909
}
