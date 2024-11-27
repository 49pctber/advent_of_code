package main

import "testing"

var test_input string = `ULL
RRDDD
LURDL
UUUUD
`

func TestGetCode(t *testing.T) {
	if have, want := GetCode(test_input), "1985"; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
}

func TestGetCode2(t *testing.T) {
	if have, want := GetCode2(test_input), "5DB3"; have != want {
		t.Fatalf("have %v, want %v", have, want)
	}
}
