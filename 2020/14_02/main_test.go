package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T) {
	want := int64(208)
	got := solveProblem("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}