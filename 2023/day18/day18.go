package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type interval_t struct {
	start int
	end   int
}

func (i interval_t) String() string {
	return fmt.Sprintf("[%d, %d]", i.start, i.end)
}

func (i *interval_t) Contains(x int) bool {
	return i.start <= x || x <= i.end
}

func (i *interval_t) Length() int {
	return i.end - i.start + 1 // endpoint inclusive
}

type intervals_t struct {
	intervals []interval_t
}

func (i intervals_t) String() string {
	s := []string{}
	for _, interval := range i.intervals {
		s = append(s, interval.String())
	}
	return strings.Join(s, " ")
}

func (i *intervals_t) Consolidate() {

	for done := false; !done; {

		done = true

		sort.SliceStable(i.intervals, func(k, j int) bool {
			return i.intervals[k].start < i.intervals[j].start
		})

		for ii := range i.intervals {
			if ii == 0 {
				continue
			}

			if i.intervals[ii-1].end >= i.intervals[ii].start {
				i.intervals[ii-1].end = max(i.intervals[ii-1].end, i.intervals[ii].end)
				i.intervals = append(i.intervals[:ii], i.intervals[ii+1:]...)
				done = false
				break
			}
		}

	}

}

func (i *intervals_t) Inside(x int) bool {

	i.Consolidate()
	for _, interval := range i.intervals {
		if interval.Contains(x) {
			return true
		}
	}

	return false
}

func (i *intervals_t) Length() int {
	i.Consolidate()
	length := 0
	for _, interval := range i.intervals {
		length += interval.Length()
	}
	return length
}

func (i intervals_t) UnionLength(j intervals_t) int {
	k := intervals_t{}
	i.Consolidate()
	j.Consolidate()
	k.intervals = append(k.intervals, i.intervals...)
	k.intervals = append(k.intervals, j.intervals...)
	return k.Length()
}

type instruction_t struct {
	direction byte
	distance  int
}

func (i instruction_t) String() string {
	return fmt.Sprintf("%c %v", i.direction, i.distance)
}

func ParseDay18Input(s string, part1 bool) []instruction_t {
	re := regexp.MustCompile(`^(\w) (\d+) \(#(\w{5})(\w)\)$`)

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	instructions := make([]instruction_t, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := re.FindStringSubmatch(scanner.Text())
		if part1 {
			distance, err := strconv.Atoi(r[2])
			if err != nil {
				log.Fatal(err)
			}
			instruction := instruction_t{direction: r[1][0], distance: distance}
			instructions = append(instructions, instruction)
		} else {
			distance, err := strconv.ParseInt(r[3], 16, 64)
			if err != nil {
				log.Fatal(err)
			}
			var direction byte
			switch r[4][0] {
			case '0':
				direction = 'R'
			case '1':
				direction = 'D'
			case '2':
				direction = 'L'
			case '3':
				direction = 'U'
			}
			instruction := instruction_t{direction: direction, distance: int(distance)}
			instructions = append(instructions, instruction)
		}
	}

	return instructions
}

func computeExtrema(instructions []instruction_t) (x_min, y_min, x_max, y_max int) {
	var x, y int

	// find number of rows/columns
	for _, instruction := range instructions {
		switch instruction.direction {
		case 'L':
			x -= instruction.distance
		case 'U':
			y -= instruction.distance
		case 'R':
			x += instruction.distance
		case 'D':
			y += instruction.distance
		default:
			log.Fatal("invalid direction")
		}

		x_min, y_min = min(x, x_min), min(y, y_min)
		x_max, y_max = max(x, x_max), max(y, y_max)
	}
	return x_min, y_min, x_max, y_max
}

func computePoi(instructions []instruction_t) map[int]map[int]struct{} {
	poi := make(map[int]map[int]struct{}) // points of interest

	x_min, y_min, _, _ := computeExtrema(instructions)
	x, y := -x_min, -y_min

	x_min, y_min, x_max, y_max := 0, 0, 0, 0
	for _, instruction := range instructions {
		switch instruction.direction {
		case 'L':
			x -= instruction.distance
		case 'U':
			y -= instruction.distance
		case 'R':
			x += instruction.distance
		case 'D':
			y += instruction.distance
		default:
			log.Fatal("invalid direction")
		}

		if _, exists := poi[y]; !exists {
			poi[y] = make(map[int]struct{})

		}
		poi[y][x] = struct{}{}

		x_min, y_min = min(x, x_min), min(y, y_min)
		x_max, y_max = max(x, x_max), max(y, y_max)
	}

	if x_min != 0 || y_min != 0 {
		log.Fatal("these should be zero")
	}

	return poi
}

func computeVolume(instructions []instruction_t) int {

	poi := computePoi(instructions) // points of interest

	coi := make(map[int]struct{}, 0) // columns of interest

	roi := make([]int, 0, len(poi)) // rows of interest
	for k := range poi {
		roi = append(roi, k)
	}
	sort.Ints(roi)

	var prev_intervals intervals_t
	var prev_row int
	var volume int

	for i, row := range roi {

		var intervals intervals_t

		cols := make([]int, 0, len(poi[row]))
		for k := range poi[row] {
			cols = append(cols, k)
		}
		sort.Ints(cols)

		if i == 0 {
			for j, col := range cols {
				coi[col] = struct{}{}
				if j%2 == 1 {
					intervals.intervals = append(intervals.intervals, interval_t{start: cols[j-1], end: cols[j]})
				}
			}
			volume += intervals.Length()
		} else {
			for _, col := range cols {
				if _, exists := coi[col]; exists {
					delete(coi, col)
				} else {
					coi[col] = struct{}{}
				}
			}

			cols := make([]int, 0, len(coi))
			for k := range coi {
				cols = append(cols, k)
			}
			sort.Ints(cols)

			// var start, end int
			for j, col := range cols {
				coi[col] = struct{}{}
				if j%2 == 1 {
					intervals.intervals = append(intervals.intervals, interval_t{start: cols[j-1], end: cols[j]})
				}
			}

			volume += (row - prev_row - 1) * prev_intervals.Length()
			volume += intervals.UnionLength(prev_intervals)
		}

		prev_row = row
		prev_intervals = intervals
	}

	volume += prev_intervals.Length()

	return volume
}

func main() {
	fmt.Println("Day 18")

	part1 := computeVolume(ParseDay18Input(`input\input18.txt`, true))
	fmt.Printf("part1: %v\n", part1) // 62500

	part2 := computeVolume(ParseDay18Input(`input\input18.txt`, false))
	fmt.Printf("part2: %v\n", part2) // 122109860712709
}
