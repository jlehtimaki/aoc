package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var nC int
	lines := readLines(input)
	for _, l := range lines {
		if l == "" {
			if nC > count {
				count = nC
			}
			nC = 0
			continue
		}
		lineInt, _ := strconv.Atoi(l)
		nC += lineInt
	}
	return count
}

func main() {
	fmt.Println(solveProblem("input.txt"))
}