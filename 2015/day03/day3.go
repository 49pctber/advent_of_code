package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type location struct {
	x int
	y int
}

func (l *location) String() string {
	return fmt.Sprintf("%dx%d", l.x, l.y)
}

func main() {
	fmt.Println("Day 3")

	file, err := os.Open(`input\input3.txt`)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	l := location{} // Santa's location
	c := make(map[string]int)

	l1 := location{} // Santa's location (year 2)
	l2 := location{} // Robo-Santa's location (year 2)
	c2 := make(map[string]int)

	c[l.String()]++
	c2[l1.String()]++
	c2[l2.String()]++

	for scanner.Scan() {
		for i, dir := range scanner.Text() {
			var ll *location // represents Santa's or Robo-Santa's location

			if i%2 == 0 {
				ll = &l1
			} else {
				ll = &l2
			}

			switch dir {
			case '>':
				l.x++
				ll.x++
			case '<':
				l.x--
				ll.x--
			case '^':
				l.y++
				ll.y++
			case 'v':
				l.y--
				ll.y--
			default:
				log.Fatal("unexpected symbol")
			}
			c[l.String()]++
			c2[ll.String()]++

		}
	}

	fmt.Println("Number of locations:", len(c))                  // 2592
	fmt.Println("Number of locations with Robo-Santa:", len(c2)) // 2360
}
