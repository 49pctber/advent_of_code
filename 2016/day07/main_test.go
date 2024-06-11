package main

import "testing"

func TestPart1(t *testing.T) {

	if have, want := ContainsABBA(``), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ContainsABBA(`aba`), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ContainsABBA(`abba`), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ContainsABBA(`zxyabba`), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ContainsABBA(`abbayxz`), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ContainsABBA(`aaaa`), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ContainsABBA(`ioxxoj`), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := ContainsABBA(`asdfgh`), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := TlsSupport(`abba[mnop]qrst`), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := TlsSupport(`ioxxoj[asdfgh]zxcvbn`), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := TlsSupport(`abcd[bddb]xyyx`), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := TlsSupport(`aaaa[qwer]tyui`), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestPart2(t *testing.T) {
	if have, want := SslSupport(`aba[bab]xyz`), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := SslSupport(`xyx[xyx]xyx`), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := SslSupport(`aaa[kek]eke`), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := SslSupport(`zazbz[bzb]cdb`), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

}
