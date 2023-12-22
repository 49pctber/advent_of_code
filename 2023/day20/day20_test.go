package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	if have, want := PressButton(`input\input20_test.txt`, 1000), 32000000; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := PressButton(`input\input20_test1.txt`, 1000), 11687500; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}
