package main

import (
	"testing"
)

func TestSolveProblem(t *testing.T) {
	want := 175594
	got := solveProblem("0,3,6")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}

//func TestSolveProblem2(t *testing.T){
//	testMap := map[int]string{
//		1: "1,3,2",
//		10: "2,1,3",
//		27: "1,2,3",
//		78: "2,3,1",
//		438: "3,2,1",
//		1836: "3,1,2",
//	}
//	for k,v := range testMap {
//		want := k
//		got := solveProblem(v)
//
//		if want != got {
//			t.Errorf("Want: %d - Got: %d", want, got)
//		}
//	}
//}