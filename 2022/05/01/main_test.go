package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T) {
	stacks := [][]string{{"N", "Z"}, {"D", "C", "M"}, {"P"}}
	want := "CMZ"
	got := solveProblem("tinput.txt", stacks)
	if want != got {
		t.Errorf("Want: %s - Got: %s", want, got)
	}
}
