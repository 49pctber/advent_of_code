package main

import "testing"

func TestDay17(t *testing.T) {
	containers := []int{20, 15, 10, 5, 5}
	target := 25

	if have, want := Start(containers, target), 4; have != want {
		t.Errorf("have %d, want %d", have, want)
	}

	if have, want := Start2(containers, target), 2; have != want {
		t.Errorf("have %d, want %d", have, want)
	}
}
