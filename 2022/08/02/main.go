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
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[0])-1; j++ {
			// 		fmt.Printf("checking number: %d \n", lines[i][j])
			c := 1
			c = c * visibleUp(j, i, lines)
			// 		fmt.Printf("final number: %d \n", c)
			c = c * visibleDown(j, i, lines)
			// 		fmt.Printf("final number: %d \n", c)
			c = c * visibleLeft(j, i, lines)
			// 		fmt.Printf("final number: %d \n", c)
			c = c * visibleRight(j, i, lines)
			// 		fmt.Printf("final number: %d \n", c)
			if c > count {
				count = c
			}
		}
	}
	fmt.Println(count)
	return count
}

func visibleUp(x int, y int, lines [][]int) int {
	c := 0
	z := y - 1
	n := lines[y][x]
	for z >= 0 {
		if lines[z][x] >= n {
			// 		fmt.Printf("up- %d \n", c+1)
			return c + 1
		}
		c++
		z--
	}
	// fmt.Printf("up- %d \n", c)
	return c
}
func visibleDown(x int, y int, lines [][]int) int {
	c := 0
	z := y + 1
	n := lines[y][x]

	for z < len(lines) {
		if lines[z][x] >= n {
			// 		fmt.Printf("down- %d \n", c+1)
			return c + 1
		}
		c++
		z++
	}
	// fmt.Printf("down- %d \n", c)
	return c
}
func visibleLeft(x int, y int, lines [][]int) int {
	c := 0
	n := lines[y][x]
	z := x - 1
	for z >= 0 {
		if lines[y][z] >= n {
			// 		fmt.Printf("left - %d \n", c+1)
			return c + 1
		}
		c++
		z--
	}
	// fmt.Printf("left - %d \n", c)
	return c
}
func visibleRight(x int, y int, lines [][]int) int {
	c := 0
	n := lines[y][x]
	z := x + 1
	for z < len(lines[0]) {
		if lines[y][z] >= n {
			// 		fmt.Printf("right- %d \n", c+1)
			return c + 1
		}
		c++
		z++
	}
	// fmt.Printf("right- %d \n", c)
	return c
}

func main() {
	fmt.Println(solveProblem("input.txt"))
}