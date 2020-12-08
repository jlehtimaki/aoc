package main

import (
	"testing"
)


func TestGetAccumalator(t *testing.T) {
	want := 8
	got := getAccumulator("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
