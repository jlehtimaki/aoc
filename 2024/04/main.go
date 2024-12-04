package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var stringA [][]string
	lineN := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringA = append(stringA, []string{})
		for _, char := range scanner.Text() {
			stringA[lineN] = append(stringA[lineN], string(char))
		}
		lineN++
	}
	return stringA
}

func checkHorizontal(x int, ar []string) int {
	var count int
	if x+3 < len(ar) {
		if ar[x+1] == "M" && ar[x+2] == "A" && ar[x+3] == "S" {
			count++
		}
	}

	if x-3 >= 0 {
		if ar[x-1] == "M" && ar[x-2] == "A" && ar[x-3] == "S" {
			count++
		}
	}

	return count
}

func checkVertical(x, y int, ar [][]string) int {
	var count int
	if y+3 < len(ar) {
		if ar[y+1][x] == "M" && ar[y+2][x] == "A" && ar[y+3][x] == "S" {
			count++
		}
	}

	if y-3 >= 0 {
		if ar[y-1][x] == "M" && ar[y-2][x] == "A" && ar[y-3][x] == "S" {
			count++
		}
	}

	return count
}

func checkDiagonal(x, y int, ar [][]string) int {
	var count int
	if x+3 < len(ar[0]) && y+3 < len(ar) {
		if ar[y+1][x+1] == "M" && ar[y+2][x+2] == "A" && ar[y+3][x+3] == "S" {
			count++
		}
	}
	if x-3 >= 0 && y-3 >= 0 {
		if ar[y-1][x-1] == "M" && ar[y-2][x-2] == "A" && ar[y-3][x-3] == "S" {
			count++
		}
	}
	if x+3 < len(ar[0]) && y-3 >= 0 {
		if ar[y-1][x+1] == "M" && ar[y-2][x+2] == "A" && ar[y-3][x+3] == "S" {
			count++
		}
	}
	if x-3 >= 0 && y+3 < len(ar) {
		if ar[y+1][x-1] == "M" && ar[y+2][x-2] == "A" && ar[y+3][x-3] == "S" {
			count++
		}
	}
	return count
}

func checkDiagonalP2(x, y int, ar [][]string) bool {
	// M.S
	// .A.
	// M.S

	if x+1 < len(ar[0]) &&
		y+1 < len(ar) &&
		x-1 >= 0 &&
		y-1 >= 0 &&
		x+1 < len(ar[0]) &&
		y+1 < len(ar) {
		x1 := fmt.Sprintf("%s%s%s", ar[y-1][x-1], ar[y][x], ar[y+1][x+1])
		x2 := fmt.Sprintf("%s%s%s", ar[y-1][x+1], ar[y][x], ar[y+1][x-1])
		if (strings.Contains(x1, "MAS") || strings.Contains(x1, "SAM")) &&
			(strings.Contains(x2, "MAS") || strings.Contains(x2, "SAM")) {
			return true
		}
	}
	return false
}

func s1(input string) int {
	var count int
	stringA := readLines(input)

	for y, yline := range stringA {
		for x, cx := range yline {
			if cx == "X" {
				count += checkHorizontal(x, yline)
				count += checkVertical(x, y, stringA)
				count += checkDiagonal(x, y, stringA)
			}
		}
	}

	return count
}

func s2(input string) int {
	var count int
	stringA := readLines(input)
	for y, yline := range stringA {
		for x, cx := range yline {
			if cx == "A" {
				if checkDiagonalP2(x, y, stringA) {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	fmt.Printf("P1: %d \n", s1("input.txt"))
	fmt.Printf("P2: %d \n", s2("input.txt"))
}
