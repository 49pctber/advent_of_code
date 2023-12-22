package main

import "testing"

func TestHash(t *testing.T) {

	if h := Hash("rn=1"); h != 30 {
		t.Errorf("invalid hash (%d instead of 30)", h)
	}
	if h := Hash("cm-"); h != 253 {
		t.Errorf("invalid hash (%d instead of 253)", h)
	}
	if h := Hash("qp=3"); h != 97 {
		t.Errorf("invalid hash (%d instead of 97)", h)
	}
	if h := Hash("cm=2"); h != 47 {
		t.Errorf("invalid hash (%d instead of 47)", h)
	}
	if h := Hash("qp-"); h != 14 {
		t.Errorf("invalid hash (%d instead of 14)", h)
	}
	if h := Hash("pc=4"); h != 180 {
		t.Errorf("invalid hash (%d instead of 180)", h)
	}
	if h := Hash("ot=9"); h != 9 {
		t.Errorf("invalid hash (%d instead of 9)", h)
	}
	if h := Hash("ab=5"); h != 197 {
		t.Errorf("invalid hash (%d instead of 197)", h)
	}
	if h := Hash("pc-"); h != 48 {
		t.Errorf("invalid hash (%d instead of 48)", h)
	}
	if h := Hash("pc=6"); h != 214 {
		t.Errorf("invalid hash (%d instead of 214)", h)
	}
	if h := Hash("ot=7"); h != 231 {
		t.Errorf("invalid hash (%d instead of 231)", h)
	}

}
