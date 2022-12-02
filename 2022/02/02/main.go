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
	// A Rock
	// B Paper
	// C Scissors
	// X LOSE
	// Y DRAW
	// Z WIN
	points := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	wins := map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}

	for _, l := range lines {
		s := strings.Split(l, " ")
		if s[1] == "Z" {
			count += 6
			for k, v := range wins {
				if v == s[0] {
					count += points[k]
					break
				}
			}
		} else if s[1] == "Y" {
			count += points[s[0]]
			count += 3
		} else {
			count += points[wins[s[0]]]
		}
	}
	return count
}

func main() {
	fmt.Println(solveProblem("input.txt"))
}