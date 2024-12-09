package main

import (
	"strings"
	"testing"
)

func TestOutputDisks(t *testing.T) {
	want := "0..111....22222"
	got := outputDisks("12345")
	if want != strings.Join(got, "") {
		t.Errorf("Want: %s - Got: %s", want, got)
	}
}

func TestOutputDisks2(t *testing.T) {
	want := "00...111...2...333.44.5555.6666.777.888899"
	disks := readLines("tinput.txt")
	got := outputDisks(disks)
	if want != strings.Join(got, "") {
		t.Errorf("Want: %s - Got: %s", want, got)
	}
}

func TestAmphiPod(t *testing.T) {
	type test struct {
		input []string
		want  string
	}

	tests := []test{
		{
			input: []string{
				"0",
				".",
				".",
				"1",
				"1",
				"1",
				".",
				".",
				".",
				".",
				"2",
				"2",
				"2",
				"2",
				"2",
			},
			want: "022111222......",
		},
		{
			input: []string{
				"0",
				"0",
				".",
				".",
				".",
				"1",
				"1",
				"1",
				".",
				".",
				".",
				"2",
				".",
				".",
				".",
				"3",
				"3",
				"3",
				".",
				"4",
				"4",
				".",
				"5",
				"5",
				"5",
				"5",
				".",
				"6",
				"6",
				"6",
				"6",
				".",
				"7",
				"7",
				"7",
				".",
				"8",
				"8",
				"8",
				"8",
				"9",
				"9",
			},
			want: "0099811188827773336446555566..............",
		},
	}

	for _, tc := range tests {
		got := strings.Join(amphiPod(tc.input), "")
		if tc.want != got {
			t.Errorf("Want: %s - Got: %s", tc.want, got)
		}
	}
}

func TestAmphiPod2(t *testing.T) {
	type test struct {
		input []string
		want  string
	}

	tests := []test{
		{
			input: []string{
				"0",
				"0",
				".",
				".",
				".",
				"1",
				"1",
				"1",
				".",
				".",
				".",
				"2",
				".",
				".",
				".",
				"3",
				"3",
				"3",
				".",
				"4",
				"4",
				".",
				"5",
				"5",
				"5",
				"5",
				".",
				"6",
				"6",
				"6",
				"6",
				".",
				"7",
				"7",
				"7",
				".",
				"8",
				"8",
				"8",
				"8",
				"9",
				"9",
			},
			want: "00992111777.44.333....5555.6666.....8888..",
		},
	}

	for _, tc := range tests {
		got := strings.Join(amphiPod2(tc.input), "")
		if tc.want != got {
			t.Errorf("Want: %s - Got: %s", tc.want, got)
		}
	}
}

func TestSolveProblem1(t *testing.T) {
	want := 1928
	got := s1("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}

func TestSolveProblem2(t *testing.T) {
	want := 2858
	got := s2("tinput.txt")
	if want != got {
		t.Errorf("Want: %d - Got: %d", want, got)
	}
}
