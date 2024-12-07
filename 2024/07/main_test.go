package main

import (
	"testing"
)

func TestSolveProblem0(t *testing.T) {
	keyValues := map[int][]int{
		190:  {0, 10, 19},
		3267: {81, 40, 27},
	}
	for k, v := range keyValues {
		if !findMatch(k, v, "p1") {
			t.Errorf("Want: true - got: false (%d:%v)", k, v)
		}
	}
}

func TestSolveProblem01(t *testing.T) {
	keyValues := map[int][]int{
		156:  {15, 6},
		7290: {6, 8, 6, 15},
		192:  {17, 8, 11},
	}

	for k, v := range keyValues {
		if !findMatch(k, v, "p2") {
			t.Errorf("Want: true - got: false (%d:%v)", k, v)
		}
	}
}

func TestSolveProblem1(t *testing.T) {
	want := 3749
	got := s1("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}

func TestSolveProblem2(t *testing.T) {
	want := 11387
	got := s2("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
