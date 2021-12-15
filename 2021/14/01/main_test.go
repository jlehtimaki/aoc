package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T) {
	want := 1588
	got := solveProblem("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
