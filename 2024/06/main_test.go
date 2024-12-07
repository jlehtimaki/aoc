package main

import (
	"testing"
)

func TestSolveProblem1(t *testing.T) {
	want := 41
	got, _ := s1("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}

func TestSolveProblem2(t *testing.T) {
	want := 6
	_, d := s1("tinput.txt")
	got := s2("tinput.txt", d)
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
