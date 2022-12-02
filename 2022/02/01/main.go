package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var stringArray []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringArray = append(stringArray, scanner.Text())
	}
	return stringArray
}

func solveProblem(input string) int {
	var count int
	lines := readLines(input)
	// X & A Rock
	// Y & B Paper
	// Z & C Scissors
	points := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	for _, l := range lines {
		s := strings.Split(l, " ")
		if s[0] == "A" && s[1] == "Y" {
			count += 6
		} else if s[0] == "B" && s[1] == "Z" {
			count += 6
		} else if s[0] == "C" && s[1] == "X" {
			count += 6
		} else if (s[0] == "A" && s[1] == "X") || (s[0] == "B" && s[1] == "Y") || (s[0] == "C" && s[1] == "Z") {
			count += 3
		}
		count += points[s[1]]
	}
	return count
}

func main() {
	fmt.Println(solveProblem("input.txt"))
}