package main

import "testing"

func TestDistance1(t *testing.T) {
	if have, want := GetPassword("abc"), "18f47a30"; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestDistance2(t *testing.T) {
	if have, want := GetPassword2("abc"), "05ace8e3"; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}
