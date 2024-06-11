package main

import "testing"

const test_input string = `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

func TestPart1(t *testing.T) {
	if have, want := Decode(test_input), "easter"; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestPart2(t *testing.T) {
	if have, want := Decode2(test_input), "advent"; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}
