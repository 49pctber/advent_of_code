package main

import "testing"

func TestPossibleArrangements(t *testing.T) {
	var input string
	var have, want int

	input = "???.### 1,1,3"
	have = possibleArrangements(input)
	want = 1
	if have != want {
		t.Errorf("%q resulted in %d, not %d", input, have, want)
	}

	input = ".??..??...?##. 1,1,3"
	have = possibleArrangements(input)
	want = 4
	if have != want {
		t.Errorf("%q resulted in %d, not %d", input, have, want)
	}

	input = "?#?#?#?#?#?#?#? 1,3,1,6"
	have = possibleArrangements(input)
	want = 1
	if have != want {
		t.Errorf("%q resulted in %d, not %d", input, have, want)
	}

	input = "????.#...#... 4,1,1"
	have = possibleArrangements(input)
	want = 1
	if have != want {
		t.Errorf("%q resulted in %d, not %d", input, have, want)
	}

	input = "????.######..#####. 1,6,5"
	have = possibleArrangements(input)
	want = 4
	if have != want {
		t.Errorf("%q resulted in %d, not %d", input, have, want)
	}

	input = "?###???????? 3,2,1"
	have = possibleArrangements(input)
	want = 10
	if have != want {
		t.Errorf("%q resulted in %d, not %d", input, have, want)
	}

}

func TestPossibleArrangements2(t *testing.T) {
	var input string
	var have, want int

	input = "???.### 1,1,3"
	have = possibleArrangements2(input)
	want = 1
	if have != want {
		t.Errorf("%q resulted in %d, not %d", input, have, want)
	}

	input = ".??..??...?##. 1,1,3"
	have = possibleArrangements2(input)
	want = 16384
	if have != want {
		t.Errorf("%q resulted in %d, not %d", input, have, want)
	}

	input = "?#?#?#?#?#?#?#? 1,3,1,6"
	have = possibleArrangements2(input)
	want = 1
	if have != want {
		t.Errorf("%q resulted in %d, not %d", input, have, want)
	}

	input = "????.#...#... 4,1,1"
	have = possibleArrangements2(input)
	want = 16
	if have != want {
		t.Errorf("%q resulted in %d, not %d", input, have, want)
	}

	input = "????.######..#####. 1,6,5"
	have = possibleArrangements2(input)
	want = 2500
	if have != want {
		t.Errorf("%q resulted in %d, not %d", input, have, want)
	}

	input = "?###???????? 3,2,1"
	have = possibleArrangements2(input)
	want = 506250
	if have != want {
		t.Errorf("%q resulted in %d, not %d", input, have, want)
	}

}
