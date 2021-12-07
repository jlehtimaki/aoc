package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T) {
	want := 26984457539
	got := solveProblem("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
