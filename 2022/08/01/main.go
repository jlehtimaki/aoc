package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var stringArray [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := []int{}
		for _, x := range scanner.Text() {
			n, _ := strconv.Atoi(string(x))
			s = append(s, n)
		}
		stringArray = append(stringArray, s)
	}
	return stringArray
}

func solveProblem(input string) int {
	var count int
	lines := readLines(input)
	count += len(lines) * 2
	count += (len(lines[0]) * 2) - 4
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[0])-1; j++ {
			fmt.Printf("checking number: %d \n", lines[i][j])
			if visibleUp(j, i, lines) {
				count++
				continue
			}
			if visibleDown(j, i, lines) {
				count++
				continue
			}
			if visibleLeft(j, i, lines) {
				count++
				continue
			}
			if visibleRight(j, i, lines) {
				count++
				continue
			}
		}
	}
	return count
}

func visibleUp(x int, y int, lines [][]int) bool {
	z := y - 1
	n := lines[y][x]
	for z >= 0 {
		if lines[z][x] >= n {
			return false
		}
		z--
	}
	return true
}
func visibleDown(x int, y int, lines [][]int) bool {
	z := y + 1
	n := lines[y][x]

	for z < len(lines) {
		if lines[z][x] >= n {
			return false
		}
		z++
	}
	return true
}
func visibleLeft(x int, y int, lines [][]int) bool {
	n := lines[y][x]
	z := x - 1
	for z >= 0 {
		if lines[y][z] >= n {
			return false
		}
		z--
	}
	return true
}
func visibleRight(x int, y int, lines [][]int) bool {
	n := lines[y][x]
	z := x + 1
	for z < len(lines[0]) {
		if lines[y][z] >= n {
			return false
		}
		z++
	}
	return true
}

func main() {
	fmt.Println(solveProblem("input.txt"))
}