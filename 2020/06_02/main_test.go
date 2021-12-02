package main


import (
	"testing"
)


func TestCountRightAnswer(t *testing.T) {
	got := countRightAnswers("test_input.txt")
	want := 6

	if got != want {
		t.Errorf("Got: %d - Want: %d", got, want)
	}
}

func TestReadLines(t *testing.T){
	got := len(readLines("test_input.txt"))
	want := 5

	if got != want {
		t.Errorf("Got: %d - Want: %d", got, want)
	}
}