package main

import (
	"bufio"
	"fmt"
	"os"
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
	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}