package main

import "testing"

func TestPart1(t *testing.T) {
	if have, want := DecompressedLength(`ADVENT`), 6; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := DecompressedLength(`A(1x5)BC`), 7; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := DecompressedLength(`(3x3)XYZ`), 9; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := DecompressedLength(`A(2x2)BCD(2x2)EFG`), 11; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := DecompressedLength(`(6x1)(1x3)A`), 6; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := DecompressedLength(`X(8x2)(3x3)ABCY`), 18; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestPart2(t *testing.T) {
	if have, want := 0, 0; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}
