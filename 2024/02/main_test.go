package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T) {
	want := 2
	got := s1("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}

func TestSolveProblemS2(t *testing.T) {
	want := 4
	got := s2("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
