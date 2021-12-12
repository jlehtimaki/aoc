package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	intArray := [][]int{}
	lines := readLines(input)
	for n, line := range lines {
		intArray = append(intArray, []int{})
		l := strings.Split(line, "")
		for _, ll := range l {
			tInt, _ := strconv.Atoi(ll)
			intArray[n] = append(intArray[n], tInt)
		}
	}


	lowPoints := map[string]int{}
	for yn, y := range intArray {
		for xn, x := range y {
			var adjacent []int
			dummy := true
			if xn == 0 {
				adjacent = append(adjacent, intArray[yn][xn+1])
			} else if xn == (len(y)-1) {
				adjacent = append(adjacent, intArray[yn][xn-1])
			} else {
				adjacent = append(adjacent, intArray[yn][xn-1])
				adjacent = append(adjacent, intArray[yn][xn+1])
			}
			if yn == 0 {
				adjacent = append(adjacent, intArray[yn+1][xn])
			} else if yn == (len(intArray)-1) {
				adjacent = append(adjacent, intArray[yn-1][xn])
			} else {
				adjacent = append(adjacent, intArray[yn-1][xn])
				adjacent = append(adjacent, intArray[yn+1][xn])
			}
			for _, a := range adjacent {
				if x >= a {
					dummy = false
					break
				}
			}
			if dummy {
				if x != 9 {
					lp := fmt.Sprintf("%d,%d", yn, xn)
					lowPoints[lp] = 0
				}
			}
		}
	}

	var points []int
	for k, _ := range lowPoints {
		points = append(points,findLowPoints(k,intArray))
		break
	}

	return count
}

func findLowPoints(coordinates string, intArray [][]int) int {
	var points int
	coordinatesArray := strings.Split(coordinates, ",")

	y,_ := strconv.Atoi(coordinatesArray[0])
	x,_ := strconv.Atoi(coordinatesArray[1])

	originalNumber := intArray[y][x]
	if  originalNumber == 9 {
		return 0
	}

	adjacent := map[string]int{}
	if x == 0 {
		coord := fmt.Sprintf("%d,%d", y, x+1)
		adjacent[coord] = intArray[y][x+1]
	} else if x == (len(intArray[y])-1) {
		coord := fmt.Sprintf("%d,%d", y, x-1)
		adjacent[coord] = intArray[y][x-1]
	} else {
		coord := fmt.Sprintf("%d,%d", y, x-1)
		adjacent[coord] = intArray[y][x-1]
		coord = fmt.Sprintf("%d,%d", y, x+1)
		adjacent[coord] = intArray[y][x+1]
	}
	if y == 0 {
		coord := fmt.Sprintf("%d,%d", y+1, x)
		adjacent[coord] = intArray[y+1][x]
	} else if y == (len(intArray)-1) {
		coord := fmt.Sprintf("%d,%d", y-1, x)
		adjacent[coord] = intArray[y-1][x]
	} else {
		coord := fmt.Sprintf("%d,%d", y-1, x)
		adjacent[coord] = intArray[y-1][x]
		coord = fmt.Sprintf("%d,%d", y+1, x)
		adjacent[coord] = intArray[y+1][x]
	}

	fmt.Println(originalNumber)
	fmt.Println(adjacent)
	for k, v := range adjacent {
		fmt.Println(v)
		if (v + 1) == originalNumber {
			points = points + 1
			points = points + findLowPoints(k, intArray)
		}
	}
	fmt.Println(points)
	

	return points
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}