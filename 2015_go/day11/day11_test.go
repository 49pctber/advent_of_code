package main

import (
	"testing"
)

func TestValidPassword(t *testing.T) {
	var input string
	var have, want bool

	input = "hijklmmn"
	have = ValidPassword(input)
	want = false
	if have != want {
		t.Errorf("%q resulted in %t, but should have been %t", input, have, want)
	}

	input = "abbceffg"
	have = ValidPassword(input)
	want = false
	if have != want {
		t.Errorf("%q resulted in %t, but should have been %t", input, have, want)
	}

	input = "abbcegjk"
	have = ValidPassword(input)
	want = false
	if have != want {
		t.Errorf("%q resulted in %t, but should have been %t", input, have, want)
	}

	input = "abcdefgh"
	have = ValidPassword(input)
	want = false
	if have != want {
		t.Errorf("%q resulted in %t, but should have been %t", input, have, want)
	}

	input = "ghijklmn"
	have = ValidPassword(input)
	want = false
	if have != want {
		t.Errorf("%q resulted in %t, but should have been %t", input, have, want)
	}

	input = "abcdffaa"
	have = ValidPassword(input)
	want = true
	if have != want {
		t.Errorf("%q resulted in %t, but should have been %t", input, have, want)
	}

	input = "ghjaabcc"
	have = ValidPassword(input)
	want = true
	if have != want {
		t.Errorf("%q resulted in %t, but should have been %t", input, have, want)
	}

}

func TestIncrementPassword(t *testing.T) {
	var input, have, want string

	input = "aaaaaaaa"
	have = incrementPassword(input)
	want = "aaaaaaab"
	if have != want {
		t.Errorf("%q resulted in %q, but should have been %q", input, have, want)
	}

	input = "zzzzzzzz"
	have = incrementPassword(input)
	want = "aaaaaaaa"
	if have != want {
		t.Errorf("%q resulted in %q, but should have been %q", input, have, want)
	}

}

func TestFindNextPassword(t *testing.T) {
	var input, have, want string

	input = "abcdefgh"
	have = FindNextPassword(input)
	want = "abcdffaa"
	if have != want {
		t.Errorf("%q resulted in %q, but should have been %q", input, have, want)
	}

	input = "ghijklmn"
	have = FindNextPassword(input)
	want = "ghjaabcc"
	if have != want {
		t.Errorf("%q resulted in %q, but should have been %q", input, have, want)
	}

}
