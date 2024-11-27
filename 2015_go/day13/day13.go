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

type person_t int

const (
	Alice person_t = iota
	Bob
	Carol
	David
	Eric
	Frank
	George
	Mallory
	Self
)

var lut map[string]person_t

const n_seats int = 8 // uncomment for part 1
// const n_seats int = 9 // uncomment for part 2
var happiness [n_seats][n_seats]int

func init() {
	lut = make(map[string]person_t)
	lut["Alice"] = Alice
	lut["Bob"] = Bob
	lut["Carol"] = Carol
	lut["David"] = David
	lut["Eric"] = Eric
	lut["Frank"] = Frank
	lut["George"] = George
	lut["Mallory"] = Mallory
	// lut["Self"] = Self // uncomment for part 2
}

type arrangement_t struct {
	seating   []person_t
	happiness int
}

func (a *arrangement_t) Compute() int {
	a.happiness = 0
	for i := 1; i <= len(a.seating); i++ {
		person1 := a.seating[i%len(a.seating)]
		person2 := a.seating[i-1]
		a.happiness += happiness[person1][person2] + happiness[person2][person1]
	}
	return a.happiness
}

func (a arrangement_t) String() string {
	return fmt.Sprintf("%v (Happiness: %d)", a.seating, a.happiness)
}

func maximize(c arrangement_t) (maximum int) {
	if len(c.seating) == n_seats {
		return c.Compute()
	}

	maximum = 0

	for _, person := range lut {
		if !slices.Contains(c.seating, person) {
			nc := arrangement_t{happiness: c.happiness}

			nc.seating = make([]person_t, len(c.seating))
			copy(nc.seating, c.seating)
			nc.seating = append(nc.seating, person)

			maximum = max(maximum, maximize(nc))
		}
	}

	return maximum
}

func main() {
	fmt.Println("Day 13")
	re := regexp.MustCompile(`(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+)\.`)

	file, err := os.Open(`input.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := re.FindStringSubmatch(scanner.Text())
		delta, err := strconv.Atoi(r[3])
		if err != nil {
			log.Fatal(err)
		}
		if r[2] == "lose" {
			delta *= -1
		}
		happiness[lut[r[1]]][lut[r[4]]] = delta
	}

	// calculate optimal seating arrangement by recursively searching orders
	maximum := maximize(arrangement_t{})
	fmt.Printf("maximum: %v\n", maximum)

}
