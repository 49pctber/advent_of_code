package main

import (
	"path/filepath"
	"testing"
)

func TestMain(t *testing.T) {
	wfs, mps, err := ParseDay19Input(filepath.Join("input", "input19.txt"))
	if err != nil {
		t.Fatal(err)
	}

	if !mps[0].Accept(wfs) {
		t.Error()
	}

	if mps[1].Accept(wfs) {
		t.Error()
	}

	if !mps[2].Accept(wfs) {
		t.Error()
	}

	if mps[3].Accept(wfs) {
		t.Error()
	}

	if !mps[4].Accept(wfs) {
		t.Error()
	}

	sum := 0
	for _, mp := range mps {
		if mp.Accept(wfs) {
			sum += mp.SumRatings()
		}
	}
	if sum != 19114 {
		t.Error()
	}

	if have, want := Combinations(wfs), 167409079868000; want != have {
		t.Errorf("want %d, have %d", want, have)
	}

}
