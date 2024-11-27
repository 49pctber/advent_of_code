package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// stores essential information about a given search
type state_t struct {
	corrupted  string
	next_index int
	n_idxs     int
}

// arrangement of hot springs
type arrangement_t struct {
	corrupted string
	groups    []int
	idxs      []int
	max_idxs  []int
}

// to store number of arrangments associated with a given state
var memo map[state_t]int

/*
Represents the state of a search. When this state is reached, the number of arragements can be stored upon return to prevent unnecessary recalculation
*/
func (a arrangement_t) GetState(ni int) state_t {
	return state_t{
		corrupted:  a.corrupted,
		next_index: ni,
		n_idxs:     len(a.idxs),
	}
}

func (a arrangement_t) String() string {
	chars := make([]string, len(a.corrupted))
	for i := 0; i < len(chars); i++ {
		chars[i] = "."
	}

	for i, n := range a.idxs {
		for j := n; j < n+a.groups[i] && j < len(chars); j++ {
			chars[j] = "#"
		}
	}
	return fmt.Sprintf("%v %v %v\n%v", a.corrupted, a.groups, a.idxs, strings.Join(chars, ""))
}

// This is used to reduce the search space by limiting the number of indexes to search over for a given group
func (a *arrangement_t) ComputeMaxIndexes() {

	a.max_idxs = make([]int, len(a.groups))
	for i := range a.groups {
		if i > 0 {
			a.max_idxs[i] = a.max_idxs[i-1] + a.groups[i-1] + 1
		} else {
			a.max_idxs[0] = 0
		}
	}

	d := len(a.corrupted) - (a.max_idxs[len(a.max_idxs)-1] + a.groups[len(a.max_idxs)-1])

	for i := len(a.corrupted) - 1; a.corrupted[i] == '.'; i-- {
		d--
	}

	for i := range a.max_idxs {
		a.max_idxs[i] += d
	}
}

func countPossibleArrangements(a arrangement_t, next_index int) int {

	// look at memos first
	state := a.GetState(next_index)
	if count, ok := memo[state]; ok {
		return count
	}

	// determine if we found a valid string
	if len(a.idxs) == len(a.groups) {
		// check that remaining are '?' or '.'
		valid := true
		for i := next_index; i < len(a.corrupted); i++ {
			if a.corrupted[i] == '#' {
				valid = false
				break
			}
		}
		if valid {
			memo[state] = 1 // record memo
			return 1
		} else {
			memo[state] = 0 // record memo
			return 0
		}
	}

	sum := 0
	group_number := len(a.idxs)

	// find next starting index
	starting_index := next_index
	for ; starting_index <= a.max_idxs[group_number]; starting_index++ {
		if a.corrupted[starting_index] != '.' {
			break
		}
	}

	for done := false; starting_index <= a.max_idxs[group_number] && !done; starting_index++ {
		if a.corrupted[starting_index] == '.' {
			continue
		} else if a.corrupted[starting_index] == '#' {
			done = true
		}

		// check that there's enough room at the end of the string
		if starting_index+a.groups[group_number]-1 >= len(a.corrupted) {
			break
		}

		// check that group consists of # or ?
		valid_group := true
		j := starting_index
		for ; j < starting_index+a.groups[group_number]; j++ {
			if a.corrupted[j] == '.' {
				valid_group = false
				break
			}
		}
		if !valid_group {
			continue
		}

		anew := arrangement_t{corrupted: a.corrupted}
		anew.groups = make([]int, len(a.groups))
		copy(anew.groups, a.groups)
		anew.idxs = make([]int, len(a.idxs))
		copy(anew.idxs, a.idxs)
		anew.max_idxs = make([]int, len(a.max_idxs))
		copy(anew.max_idxs, a.max_idxs)
		anew.idxs = append(anew.idxs, starting_index)

		if j+1 <= len(anew.corrupted) {
			if a.corrupted[j] != '#' {
				sum += countPossibleArrangements(anew, j+1)
			}
		} else {
			sum += countPossibleArrangements(anew, j+1)
		}

	}
	memo[state_t{corrupted: a.corrupted, next_index: starting_index, n_idxs: (len(a.idxs))}] = sum
	memo[state] = sum
	return sum
}

func possibleArrangements(s string) int {
	re := regexp.MustCompile(`^([?#\.]*) (.+)$`)
	r := re.FindStringSubmatch(s)
	a := arrangement_t{corrupted: r[1]}

	group_str := strings.Split(r[2], ",")
	for _, nstr := range group_str {
		n, err := strconv.Atoi(nstr)
		if err != nil {
			log.Fatal(err)
		}
		a.groups = append(a.groups, n)
	}

	a.ComputeMaxIndexes()

	return countPossibleArrangements(a, 0)
}

func possibleArrangements2(s string) int {
	re := regexp.MustCompile(`^([?#\.]*) (.+)$`)
	r := re.FindStringSubmatch(s)
	a := arrangement_t{corrupted: strings.Join([]string{r[1], r[1], r[1], r[1], r[1]}, ",")}

	group_str := strings.Split(r[2], ",")
	groups := []int{}
	for _, nstr := range group_str {
		n, err := strconv.Atoi(nstr)
		if err != nil {
			log.Fatal(err)
		}
		groups = append(groups, n)
	}

	a.groups = append(a.groups, groups...)
	a.groups = append(a.groups, groups...)
	a.groups = append(a.groups, groups...)
	a.groups = append(a.groups, groups...)
	a.groups = append(a.groups, groups...)

	a.ComputeMaxIndexes()

	return countPossibleArrangements(a, 0)
}

func init() {
	memo = make(map[state_t]int)
}

func main() {
	fmt.Println("Day 12")

	file, err := os.Open(filepath.Join("input", "input12.txt"))
	if err != nil {
		log.Fatal(err)
	}

	sum := 0  // part 1
	sum2 := 0 // part 2

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// clear unneeded memos
		memo = make(map[state_t]int)

		input := scanner.Text()
		sum += possibleArrangements(input)
		sum2 += possibleArrangements2(scanner.Text())
	}

	fmt.Printf("sum: %v\n", sum)   // 6949
	fmt.Printf("sum2: %v\n", sum2) // 51456609952403
}
