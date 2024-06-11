package main

import "testing"

var test_input string = `5-8
0-2
4-7`

func TestPart1(t *testing.T) {
	if have, want := LowestAllowed(test_input), uint32(3); have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}
