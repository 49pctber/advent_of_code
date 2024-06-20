package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := InputString("test.txt")
	if have, want := GetBot(input), 2; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestPart2(t *testing.T) {
	if have, want := 0, 0; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}
