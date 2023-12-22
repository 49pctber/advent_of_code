package main

import (
	"testing"
)

func TestLookAndSay(t *testing.T) {
	var have, want string

	have = "1"
	want = "11"
	if LookAndSay(have) != want {
		t.Errorf("%s resulted in %s, but should have been %s", have, LookAndSay(have), want)
	}

	have = "11"
	want = "21"
	if LookAndSay(have) != want {
		t.Errorf("%s resulted in %s, but should have been %s", have, LookAndSay(have), want)
	}

	have = "21"
	want = "1211"
	if LookAndSay(have) != want {
		t.Errorf("%s resulted in %s, but should have been %s", have, LookAndSay(have), want)
	}

	have = "1211"
	want = "111221"
	if LookAndSay(have) != want {
		t.Errorf("%s resulted in %s, but should have been %s", have, LookAndSay(have), want)
	}

	have = "111221"
	want = "312211"
	if LookAndSay(have) != want {
		t.Errorf("%s resulted in %s, but should have been %s", have, LookAndSay(have), want)
	}

}
