package main

import (
	"testing"
)

func TestNice(t *testing.T) {
	var s string

	s = "ugknbfddgicrmopn"
	if Nice(s) != true {
		t.Errorf("%s is nice, but it was claimed to be naughty", s)
	}

	s = "aaa"
	if Nice(s) != true {
		t.Errorf("%s is nice, but it was claimed to be naughty", s)
	}

	s = "jchzalrnumimnmhp"
	if Nice(s) == true {
		t.Errorf("%s is naughty, but it was claimed to be nice", s)
	}

	s = "haegwjzuvuyypxyu"
	if Nice(s) == true {
		t.Errorf("%s is naughty, but it was claimed to be nice", s)
	}

	s = "dvszwmarrgswjxmb"
	if Nice(s) == true {
		t.Errorf("%s is naughty, but it was claimed to be nice", s)
	}

}

func TestNice2(t *testing.T) {
	var s string

	s = "qjhvhtzxzqqjkmpb"
	if Nice2(s) != true {
		t.Errorf("%s is nice, but it was claimed to be naughty", s)
	}

	s = "xxyxx"
	if Nice2(s) != true {
		t.Errorf("%s is nice, but it was claimed to be naughty", s)
	}

	s = "uurcxstgmygtbstg"
	if Nice2(s) == true {
		t.Errorf("%s is naughty, but it was claimed to be nice", s)
	}

	s = "ieodomkazucvgmuy"
	if Nice2(s) == true {
		t.Errorf("%s is naughty, but it was claimed to be nice", s)
	}

}
