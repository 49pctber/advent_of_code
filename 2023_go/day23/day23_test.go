package main

import (
	"fmt"
	"testing"
)

func TestDay23(t *testing.T) {
	maze := parseInput(`input\input_test.txt`, false)

	for start := range maze.forks {
		for end, distance := range maze.forks[start] {
			fmt.Printf("%d -> %d (%d)\n", start, end, distance)
		}
	}

	if have, want := maze.longestHike(), 94; have != want {
		t.Errorf("have %d, want %d", have, want)
	}

	maze = parseInput(`input\input_test.txt`, true)

	if have, want := maze.longestHike(), 154; have != want {
		t.Errorf("have %d, want %d", have, want)
	}
}
