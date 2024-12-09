package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	y int
	x int
}

func readLines(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var stringArray [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		stringArray = append(stringArray, chars)
	}
	return stringArray
}

func getCoordinates(input [][]string) map[string][]Pair {
	cords := make(map[string][]Pair)
	for y, line := range input {
		for x, c := range line {
			if c != "." {
				cords[c] = append(cords[c], Pair{x, y})
			}
		}
	}
	return cords
}

func findAntinodesP1(pairs []Pair, maxY, maxX int) []Pair {
	antinodes := []Pair{}
	for i, pair := range pairs {
		for j := i + 1; j < len(pairs); j++ {
			xx := pair.x
			xy := pair.y
			yx := pairs[j].x
			yy := pairs[j].y

			cx, cy := xx-(yx-xx), xy-(yy-xy)
			dx, dy := yx+(yx-xx), yy+(yy-xy)

			if cx >= 0 && cx < maxX && cy >= 0 && cy < maxY {
				antinodes = append(antinodes, Pair{cx, cy})
			}
			if dx >= 0 && dx < maxX && dy >= 0 && dy < maxY {
				antinodes = append(antinodes, Pair{dx, dy})
			}
		}
	}
	return antinodes
}

func insideGrid(x, y, maxX, maxY int) bool {
	return x >= 0 && x < maxX && y >= 0 && y < maxY
}

func findAntinodesP2(pairs []Pair, maxY, maxX int) []Pair {
	var antinodes []Pair
	// always a tupple
	xx := pairs[0].x
	xy := pairs[0].y
	yx := pairs[1].x
	yy := pairs[1].y

	dx, dy := (yx - xx), (yy - xy)

	j := 0
	for {
		if insideGrid(xx-dx*j, xy-dy*j, maxX, maxY) {
			antinodes = append(antinodes, Pair{xx - dx*j, xy - dy*j})
		} else {
			break
		}
		j++
	}
	j = 0
	for {
		if insideGrid(yx+dx*j, yy+dy*j, maxX, maxY) {
			antinodes = append(antinodes, Pair{yx + dx*j, yy + dy*j})
		} else {
			break
		}
		j++
	}
	return antinodes
}

func generateCombinations(items []Pair) [][]Pair {
	var results [][]Pair
	n := len(items)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			results = append(results, []Pair{items[i], items[j]})
		}
	}
	return results
}

func s1(input string) int {
	antinodes := make(map[Pair]bool)
	lines := readLines(input)
	coords := getCoordinates(lines)
	for _, c := range coords {
		combs := generateCombinations(c)
		for _, comb := range combs {
			for _, p := range findAntinodesP1(comb, len(lines[0]), len(lines)) {
				antinodes[p] = true
			}
		}
	}
	return len(antinodes)
}

func s2(input string) int {
	antinodes := make(map[Pair]bool)
	lines := readLines(input)
	coords := getCoordinates(lines)
	for _, c := range coords {
		combs := generateCombinations(c)
		for _, comb := range combs {
			for _, p := range findAntinodesP2(comb, len(lines[0]), len(lines)) {
				antinodes[p] = true
			}
		}
	}
	return len(antinodes)
}

func main() {
	fmt.Printf("P1: %d \n", s1("input.txt"))
	fmt.Printf("P2: %d \n", s2("input.txt"))
}
