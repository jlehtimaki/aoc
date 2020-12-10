package main

import (
	"testing"
)


func TestSolveProblem(t *testing.T) {
	want := 62
	got := solveProblem("tinput.txt", 127)
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
