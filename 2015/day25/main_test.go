package main

import "testing"

func TestIndexAt(t *testing.T) {
	if have, want := indexAt(1, 1), 1; have != want {
		t.Errorf("have %d, want %d", have, want)
	}

	if have, want := indexAt(5, 2), 17; have != want {
		t.Errorf("have %d, want %d", have, want)
	}
}

func TestValAt(t *testing.T) {
	if have, want := valAt(1, 1), int64(20151125); have != want {
		t.Errorf("have %d, want %d", have, want)
	}

	if have, want := valAt(5, 2), int64(17552253); have != want {
		t.Errorf("have %d, want %d", have, want)
	}

	if have, want := valAt(6, 6), int64(27995004); have != want {
		t.Errorf("have %d, want %d", have, want)
	}
}
