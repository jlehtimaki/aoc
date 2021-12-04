package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T) {
	b := BingoCards{}
	want := 4512
	got := b.solveProblem("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
