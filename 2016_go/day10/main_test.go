package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := InputString("test.txt")
	SetSearchParamters(2, 5)
	GetBot(input)
}
