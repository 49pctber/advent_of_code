package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {

	input := `rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1
`
	instructions := strings.Split(input, "\n")

	s, err := ParseInstructions(3, 7, instructions)
	if err != nil {
		t.Error(err)
	}

	if !(s.pixels[0][1] && s.pixels[0][4] && s.pixels[0][6] && s.pixels[1][0] && s.pixels[1][2] && s.pixels[2][1]) {
		t.Error("pixels that should be lit aren't")
	}

	if have, want := s.CountLit(), 6; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestPart2(t *testing.T) {
	if have, want := 0, 0; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}
