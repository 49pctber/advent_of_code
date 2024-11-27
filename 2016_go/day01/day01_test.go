package main

import "testing"

func TestDistance1(t *testing.T) {

	if have, want := Distance1("R2, L3"), 5; have != want {
		t.Errorf("have %d, want %d", have, want)
	}

	if have, want := Distance1("R2, R2, R2"), 2; have != want {
		t.Errorf("have %d, want %d", have, want)
	}

	if have, want := Distance1("R5, L5, R5, R3"), 12; have != want {
		t.Errorf("have %d, want %d", have, want)
	}

}

func TestDistance2(t *testing.T) {

	if have, want := Distance2("R8, R4, R4, R8"), 4; have != want {
		t.Errorf("have %d, want %d", have, want)
	}

}
