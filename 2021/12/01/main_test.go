package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T) {
	want := 10
	got := solveProblem("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}

func TestSolveProblem2(t *testing.T) {
	want := 19
	got := solveProblem("tinput2.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}

func TestSolveProblem3(t *testing.T) {
	want := 226
	got := solveProblem("tinput3.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
