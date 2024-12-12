package main

import (
	"math"
	"testing"
)

func TestMain(t *testing.T) {
	if have, want := day17Search(0, 3, math.MaxInt, "test1.txt"), 102; have != want {
		t.Errorf("returned %d, wanted %d", have, want)
	}

	if have, want := day17Search(4, 10, math.MaxInt, "test1.txt"), 94; have != want {
		t.Errorf("returned %d, wanted %d", have, want)
	}

	if have, want := day17Search(4, 10, math.MaxInt, "test2.txt"), 71; have != want {
		t.Errorf("returned %d, wanted %d", have, want)
	}
}
