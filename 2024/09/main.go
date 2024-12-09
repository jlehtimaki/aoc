package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func readLines(path string) string {
	var str string
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str += scanner.Text()
	}
	return str
}

func outputDisks(str string) []string {
	var strA []string

	isFile := true
	id := 0
	for _, c := range str {
		cInt, _ := strconv.Atoi(string(c))
		ids := strconv.Itoa(id)
		if isFile {
			for range cInt {
				strA = append(strA, ids)
			}
			id++
			isFile = false
			continue
		}
		for range cInt {
			strA = append(strA, ".")
		}
		isFile = true
	}

	return strA
}

func amphiPod(amphiArray []string) []string {
	for i, v := range amphiArray {
		if v == "." {
			for j := len(amphiArray) - 1; j > i; j-- {
				if amphiArray[j] != "." {
					amphiArray[i] = amphiArray[j]
					amphiArray[j] = "."
					break
				}
			}
		}
	}

	return amphiArray
}

func addToMatrix(matrix [][]string, idx int, newSlice []string) [][]string {
	if idx < 0 || idx > len(matrix) {
		return matrix
	}

	matrix = append(matrix[:idx], append([][]string{newSlice}, matrix[idx:]...)...)
	return matrix
}

func buildSlice(n int, c string) []string {
	var slice []string
	for range n {
		slice = append(slice, c)
	}
	return slice
}

func amphiPod2(amphiArray []string) []string {
	var amphiGrid [][]string
	var tmpArray []string
	var newAmphiArray []string
	for i := 0; i < len(amphiArray); i++ {
		tmpArray = append(tmpArray, amphiArray[i])
		if i == len(amphiArray)-1 {
			amphiGrid = append(amphiGrid, tmpArray)
			continue
		}
		if amphiArray[i+1] != amphiArray[i] {
			amphiGrid = append(amphiGrid, tmpArray)
			tmpArray = nil
		}
	}

	for i := len(amphiGrid) - 1; i > 0; i-- {
		if slices.Contains(amphiGrid[i], ".") {
			continue
		}
		for j := 0; j < i; j++ {
			if slices.Contains(amphiGrid[j], ".") {
				// if good change, found slot
				if len(amphiGrid[j]) >= len(amphiGrid[i]) {
					diff := len(amphiGrid[j]) - len(amphiGrid[i])
					if diff == 0 {
						amphiGrid[i], amphiGrid[j] = amphiGrid[j], amphiGrid[i]
					} else {
						amphiGrid[j], amphiGrid[i] = amphiGrid[i], buildSlice(len(amphiGrid[i]), ".")
						amphiGrid = addToMatrix(amphiGrid, j+1, buildSlice(diff, "."))
						i++
					}
					break
				}
			}
		}
	}

	for _, v := range amphiGrid {
		newAmphiArray = append(newAmphiArray, v...)
	}
	return newAmphiArray
}

func checkSum(amph []string) int {
	var count int
	for i, v := range amph {
		if v == "." {
			continue
		}
		vInt, _ := strconv.Atoi(v)
		count += i * vInt
	}

	return count
}

func s1(input string) int {
	lines := readLines(input)
	outPut := outputDisks(lines)
	amphipod := amphiPod(outPut)
	return checkSum(amphipod)
}

func s2(input string) int {
	lines := readLines(input)
	outPut := outputDisks(lines)
	amphipod := amphiPod2(outPut)
	return checkSum(amphipod)
}

func main() {
	fmt.Printf("P1: %d \n", s1("input.txt"))
	fmt.Printf("P2: %d \n", s2("input.txt"))
}
