package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var array []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineInt,_ := strconv.Atoi(scanner.Text())
		array = append(array, lineInt)
	}
	return array
}

func checkSum(input []int, sum int) bool {
	for _, i := range input {
		for _, x := range input {
			if i != x {
				if (i + x) == sum {
					return true
				}
			}
		}
	}
	return false
}

func solveProblem(input string, number int) int {
	lines := readLines(input)
	preAmp := number
	for n, line := range lines {
		if n > preAmp {
			preArray := lines[n-preAmp:n]
			if checkSum(preArray, line) {
				continue
			}
			return line
		}
	}

	return 0
}

func main()  {
	fmt.Println(solveProblem("input.txt", 25))
}
