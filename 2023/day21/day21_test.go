package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	search := ParseDay21Input(`input21_test1.txt`, 6)

	if start_row, start_col, err := search.FindStart(); start_row != 5 || start_col != 5 || err != nil {
		t.Errorf("incorrect starting point (%d, %d) [%v]", start_row, start_col, err)
	}

	if have, want := CountGardenPlots(`input21_test1.txt`, 6), 16; have != want {
		t.Errorf("have %d, want %d", have, want)
	}
}
