package main

import "testing"

func Testmaint *testing.T) {
	comet := reindeer_t{speed: 14, fly_dur: 10, rest_dur: 127}
	// dancer := reindeer_t{speed: 16, fly_dur: 11, rest_dur: 162}

	if dist, want := comet.Distance(1), 14; dist != want {
		t.Errorf("comet 1: have %d want %d", dist, want)
	}

	if dist, want := comet.Distance(10), 140; dist != want {
		t.Errorf("comet 10: have %d want %d", dist, want)
	}

	if dist, want := comet.Distance(11), 140; dist != want {
		t.Errorf("comet 11: have %d want %d", dist, want)
	}

	if dist, want := comet.Distance(1000), 1120; dist != want {
		t.Errorf("comet 1000: have %d want %d", dist, want)
	}
}
