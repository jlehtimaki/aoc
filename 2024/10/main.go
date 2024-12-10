package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var intMatrix [][]int
	var stringArray [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringArray = append(stringArray, strings.Split(scanner.Text(), ""))
	}

	for _, y := range stringArray {
		var intArray []int

		for _, x := range y {
			xInt, _ := strconv.Atoi(x)
			intArray = append(intArray, xInt)
		}
		intMatrix = append(intMatrix, intArray)
		intArray = nil
	}

	return intMatrix
}

func findPath(x, y int, matrix [][]int, visited map[string]bool) int {
	var count int
	visited[strconv.Itoa(y)+strconv.Itoa(x)] = true
	if matrix[y][x] == 9 {
		return 1
	}
	// left
	if x > 0 {
		if matrix[y][x-1] == matrix[y][x]+1 {
			if !visited[strconv.Itoa(y)+strconv.Itoa(x-1)] {
				count += findPath(x-1, y, matrix, visited)
			}
		}
	}
	// right
	if x < len(matrix[y])-1 {
		if matrix[y][x+1] == matrix[y][x]+1 {
			if !visited[strconv.Itoa(y)+strconv.Itoa(x+1)] {
				count += findPath(x+1, y, matrix, visited)
			}
		}
	}
	// up
	if y > 0 {
		if matrix[y-1][x] == matrix[y][x]+1 {
			if !visited[strconv.Itoa(y-1)+strconv.Itoa(x)] {
				count += findPath(x, y-1, matrix, visited)
			}
		}
	}
	// down
	if y < len(matrix)-1 {
		if matrix[y+1][x] == matrix[y][x]+1 {
			if !visited[strconv.Itoa(y+1)+strconv.Itoa(x)] {
				count += findPath(x, y+1, matrix, visited)
			}
		}
	}

	return count
}

func findPathMap(
	x, y int,
	matrix [][]int,
	visited map[string]bool,
	pointMap map[string]int,
) map[string]int {
	point := strconv.Itoa(y) + strconv.Itoa(x)
	visited[point] = true
	defer func() {
		// Unmark visited before returning, enabling other paths to use this cell
		visited[point] = false
	}()

	if matrix[y][x] == 9 {
		pointMap[point] += 1
		return pointMap
	}

	// left
	if x > 0 {
		if matrix[y][x-1] == matrix[y][x]+1 {
			if !visited[strconv.Itoa(y)+strconv.Itoa(x-1)] {
				pointMap = findPathMap(x-1, y, matrix, visited, pointMap)
			}
		}
	}
	// right
	if x < len(matrix[y])-1 {
		if matrix[y][x+1] == matrix[y][x]+1 {
			if !visited[strconv.Itoa(y)+strconv.Itoa(x+1)] {
				pointMap = findPathMap(x+1, y, matrix, visited, pointMap)
			}
		}
	}
	// up
	if y > 0 {
		if matrix[y-1][x] == matrix[y][x]+1 {
			if !visited[strconv.Itoa(y-1)+strconv.Itoa(x)] {
				pointMap = findPathMap(x, y-1, matrix, visited, pointMap)
			}
		}
	}
	// down
	if y < len(matrix)-1 {
		if matrix[y+1][x] == matrix[y][x]+1 {
			if !visited[strconv.Itoa(y+1)+strconv.Itoa(x)] {
				pointMap = findPathMap(x, y+1, matrix, visited, pointMap)
			}
		}
	}

	return pointMap
}

func countMapRating(trailheads map[string]int) int {
	var count int
	for _, tr := range trailheads {
		count += tr
	}
	return count
}

func s1(input string) int {
	var count int
	matrix := readLines(input)
	for y, yx := range matrix {
		for x, xx := range yx {
			if xx == 0 {
				count += findPath(x, y, matrix, make(map[string]bool))
			}
		}
	}
	return count
}

func s2(input string) int {
	var count int
	matrix := readLines(input)
	for y, yx := range matrix {
		for x, xx := range yx {
			if xx == 0 {
				startMap := make(map[string]map[string]int)
				cord := strconv.Itoa(y) + strconv.Itoa(x)
				startMap[cord] = findPathMap(
					x,
					y,
					matrix,
					make(map[string]bool),
					make(map[string]int),
				)
				count += countMapRating(startMap[cord])
			}
		}
	}
	return count
}

func main() {
	fmt.Printf("P1: %d \n", s1("input.txt"))
	fmt.Printf("P2: %d \n", s2("input.txt"))
}
