package main

import (
	"testing"
)

func TestSolveProblem1(t *testing.T) {
	want := 18
	got := s1("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}

func TestSolveProblem2(t *testing.T) {
	want := 9
	got := s2("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
