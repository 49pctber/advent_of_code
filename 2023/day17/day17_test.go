package main

import (
	"math"
	"testing"
)

func TestMain(t *testing.T) {
	if have, want := day17Search(0, 3, math.MaxInt, `input\input17_test1.txt`), 102; have != want {
		t.Errorf("returned %d, wanted %d", have, want)
	}
}
