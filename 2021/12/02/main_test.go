package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T) {
	want := 36
	got := solveProblem("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}

//func TestSolveProblem2(t *testing.T) {
//	want := 103
//	got := solveProblem("tinput2.txt")
//	if want != got {
//		t.Errorf("Want: %d - Got: %d", want, got)
//	}
//}
//
//func TestSolveProblem3(t *testing.T) {
//	want := 3509
//	got := solveProblem("tinput3.txt")
//	if want != got {
//		t.Errorf("Want: %d - Got: %d", want, got)
//	}
//}
