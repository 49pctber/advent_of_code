package main

import (
	"testing"
)

func Testmaint *testing.T) {
	instructions := ParseDay18Input(`input\input18_test1.txt`, true)
	if have, want := computeVolume(instructions), 62; have != want {
		t.Errorf("Returned %d, not %d", have, want)
	}

	instructions = ParseDay18Input(`input\input18_test1.txt`, false)
	if have, want := computeVolume(instructions), 952408144115; have != want {
		t.Errorf("Returned %d, not %d", have, want)
	}
}

func TestIntervals(t *testing.T) {
	var interval interval_t

	interval = interval_t{start: 0, end: 10}
	if interval.Length() != 11 {
		t.Error("This should be 11")
	}

	var intervals intervals_t
	intervals.intervals = []interval_t{{start: 0, end: 10}, {start: 20, end: 30}}
	if intervals.Length() != 22 {
		t.Errorf("This should be 22 (%v)", intervals)
	}

	intervals.intervals = []interval_t{{start: -20, end: 30}, {start: 2, end: 10}}
	if intervals.Length() != 51 {
		t.Errorf("This should be 51, not %d (%v)", intervals.Length(), intervals)
	}

	intervals.intervals = []interval_t{{start: 0, end: 10}, {start: 0, end: 30}}
	if intervals.Length() != 31 {
		t.Errorf("This should be 22 (%v)", intervals)
	}

	if intervals.UnionLength(intervals_t{intervals: []interval_t{{start: 25, end: 40}}}) != 41 {
		t.Errorf("This should be 41, not %d", intervals.UnionLength(intervals_t{intervals: []interval_t{{start: 25, end: 40}}}))
	}
}
