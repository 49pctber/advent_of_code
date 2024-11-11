package main

import (
	"path/filepath"
	"testing"
)

func TestMain(t *testing.T) {
	if have, want := PressButton(filepath.Join("input", "input20.txt"), 1000), 32000000; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := PressButton(filepath.Join("input", "input20.txt"), 1000), 11687500; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}
