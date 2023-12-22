package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	test_fname := `input\input15_test.txt`

	cookie := cookie_t{ingredients: ParseDay15Input(test_fname)}
	cookie.quantities = make([]int, len(cookie.ingredients))
	for i, ingredient := range cookie.ingredients {
		switch ingredient.label {
		case "Butterscotch":
			cookie.quantities[i] = 44
		case "Cinnamon":
			cookie.quantities[i] = 56
		}
	}

	if have, want := cookie.Score(), 62842880; have != want {
		t.Errorf("have %d, want %d", have, want)
	}

	// if have, want := OptimizeCookie(test_fname), 62842880; have != want {
	// 	t.Errorf("have %d, want %d", have, want)
	// }
}
