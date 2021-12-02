package main

import (
	"testing"
)


func TestGetBags(t *testing.T) {
	got := getBags("test_input.txt")
	want := 4

	if got != want {
		t.Errorf("Got: %d - Want: %d", got, want)
	}
}

func TestReadLines(t *testing.T){
	got := len(readLines("test_input.txt"))
	want := 9

	if got != want {
		t.Errorf("Got: %d - Want: %d", got, want)
	}
}

func TestMainReadLines(t *testing.T){
	got := len(readLines("input.txt"))
	want := 594

	if got != want {
		t.Errorf("Got: %d - Want: %d", got, want)
	}
}