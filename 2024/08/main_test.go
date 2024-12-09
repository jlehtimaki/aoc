package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCoordinates(t *testing.T) {
	input := `..........
..........
..........
....a.....
..........
.....a....
..........
..........
..........
..........`
	var stringArray [][]string
	for _, line := range strings.Split(input, "\n") {
		lineSplit := strings.Split(string(line), "")
		stringArray = append(stringArray, lineSplit)
	}

	want := map[string][]Pair{
		"a": {{4, 3}, {5, 5}},
	}

	got := getCoordinates(stringArray)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %v - Got: %v", want, got)
	}
}

func TestCoordinates2(t *testing.T) {
	input := `..........
..........
..........
....a.....
........a.
.....a....
..........
..........
..........
..........`
	var stringArray [][]string
	for _, line := range strings.Split(input, "\n") {
		lineSplit := strings.Split(string(line), "")
		stringArray = append(stringArray, lineSplit)
	}

	want := map[string][]Pair{
		"a": {{4, 3}, {8, 4}, {5, 5}},
	}

	got := getCoordinates(stringArray)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %v - Got: %v", want, got)
	}
}

func TestAntiNodes(t *testing.T) {
	type test struct {
		input map[string][]Pair
		want  []Pair
	}

	tests := []test{
		{
			input: map[string][]Pair{
				"a": {{4, 3}, {5, 5}},
			},
			want: []Pair{{1, 3}, {7, 6}},
		},
		{
			input: map[string][]Pair{
				"a": {{4, 3}, {8, 4}, {5, 5}},
			},
			want: []Pair{{2, 0}, {1, 3}, {7, 6}, {6, 2}},
		},
	}

	for _, tc := range tests {
		got := findAntinodesP1(tc.input["a"], 9, 9)

		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("Want: %v - Got: %v", tc.want, got)
		}
	}
}

func TestInput2(t *testing.T) {
	input := `..........
..........
..........
....a.....
........a.
.....a....
..........
..........
..........
..........`
	var stringArray [][]string
	for _, line := range strings.Split(input, "\n") {
		lineSplit := strings.Split(string(line), "")
		stringArray = append(stringArray, lineSplit)
	}
	want := 4
	got := 0
	coords := getCoordinates(stringArray)
	for _, c := range coords {
		got += len(findAntinodesP1(c, len(stringArray[0]), len(stringArray)))
	}
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}

func TestSolveProblem1(t *testing.T) {
	want := 14
	got := s1("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}

func TestSolveProblem2(t *testing.T) {
	want := 34
	got := s2("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
