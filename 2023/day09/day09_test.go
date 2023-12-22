package main

import "testing"

func TestDay9_1(t *testing.T) {
	var have []int
	var want int

	have = []int{1, 3, 6, 10, 15, 21}
	want = 28

	if Predict(have) != want {
		t.Errorf("%v resulted in %d (should be %d)", have, Predict(have), want)
	}

	have = []int{0, 3, 6, 9, 12, 15}
	want = 18

	if Predict(have) != want {
		t.Errorf("%v resulted in %d (should be %d)", have, Predict(have), want)
	}

	have = []int{10, 13, 16, 21, 30, 45}
	want = 68

	if Predict(have) != want {
		t.Errorf("%v resulted in %d (should be %d)", have, Predict(have), want)
	}
}

func TestDay9_2(t *testing.T) {
	var have []int
	var want int

	have = []int{10, 13, 16, 21, 30, 45}
	want = 5

	if Postdict(have) != want {
		t.Errorf("%v resulted in %d (should be %d)", have, Postdict(have), want)
	}

}
