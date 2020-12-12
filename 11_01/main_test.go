package main

import (
	"strings"
	"testing"
)

func TestGetSeatOrder(t *testing.T) {
	wantRaw := []string{"#.##.##.##","#######.##","#.#.#..#..","####.##.##","#.##.##.##","#.#####.##","..#.#.....","##########","#.######.#","#.#####.##"}
	lines := readLines("tinput.txt")
	gotRaw,_ := getSeatOrder(lines)

	want := strings.Join(wantRaw, ",")
	got := strings.Join(gotRaw, ",")
	if want != got {
		t.Errorf("Want: %s - Got: %s", want, got)
	}
}

func TestGetSeatOrder2(t *testing.T) {
	line := []string{"#.##.##.##", "#######.##", "#.#.#..#..", "####.##.##", "#.##.##.##", "#.#####.##", "..#.#.....", "##########", "#.######.#", "#.#####.##"}
	wantRaw := []string{"#.LL.L#.##",
		"#LLLLLL.L#",
		"L.L.L..L..",
		"#LLL.LL.L#",
		"#.LL.LL.LL",
		"#.LLLL#.##",
		"..L.L.....",
		"#LLLLLLLL#",
		"#.LLLLLL.L",
		"#.#LLLL.##",
	}
	gotRaw, _ := getSeatOrder(line)
	want := strings.Join(wantRaw, ",")
	got := strings.Join(gotRaw, ",")
	if want != got {
		t.Errorf("Want: %s - Got: %s", want, got)
	}
}

//func TestSolveProblem(t *testing.T) {
//	lines := readLines("tinput.txt")
//	want := 37
//	got := solveProblem(lines)
//	if want != got {
//		t.Errorf("Want: %d - Got: %d", want, got)
//	}
//}
