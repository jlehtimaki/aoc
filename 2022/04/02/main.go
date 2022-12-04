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
	lines := readLines(input)
	for _, l := range lines {
		pairs := strings.Split(l, ",")
		p1 := strings.Split(pairs[0], "-")
		p2 := strings.Split(pairs[1], "-")
		p11, _ := strconv.Atoi(p1[0])
		p12, _ := strconv.Atoi(p1[1])
		p21, _ := strconv.Atoi(p2[0])
		p22, _ := strconv.Atoi(p2[1])
		if !(p12 < p21 || p22 < p11) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(solveProblem("input.txt"))
}