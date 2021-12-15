package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T) {
	want := 2188189693529
	got := solveProblem("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
