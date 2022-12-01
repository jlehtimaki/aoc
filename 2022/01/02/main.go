package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var nC int
	elfArray := []int{}
	lines := readLines(input)
	for _, l := range lines {
		if l == "" {
			elfArray = append(elfArray, nC)
			nC = 0
			continue
		}
		lineInt, _ := strconv.Atoi(l)
		nC += lineInt
	}
	elfArray = append(elfArray, nC)
	sort.Ints(elfArray)
	return elfArray[len(elfArray)-1] + elfArray[len(elfArray)-2] + elfArray[len(elfArray)-3]
}

func main() {
	fmt.Println(solveProblem("input.txt"))
}