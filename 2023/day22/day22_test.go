package main

import (
	"testing"
)

func Testmaint *testing.T) {
	bricks := ParseDay22Input(`input\input22_test.txt`)
	// fmt.Printf("bricks: %v\n", bricks)
	// n_falling_bricks := bricks.Fall()
	// fmt.Printf("bricks: %v\n", bricks)
	// fmt.Printf("n_falling_bricks: %v\n", n_falling_bricks)

	if have, want := bricks.CountCandidateBricks(), 5; have != want {
		t.Errorf("have %d, want %d", have, want)
	}

	if have, want := bricks.CountFallingBricks(), 7; have != want {
		t.Errorf("have %d, want %d", have, want)
	}

}
