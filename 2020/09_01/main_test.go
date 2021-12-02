package main

import (
	"testing"
)


func TestSolveProblem(t *testing.T) {
	want := 127
	got := solveProblem("tinput.txt", 5)
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
