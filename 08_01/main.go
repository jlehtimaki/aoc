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

func getAccumulator(input string) int{

	accumulator := 0

	lines := readLines(input)
	lineNumber := 0
	usedLines := map[int]int{}

	for true {
		if _, ok := usedLines[lineNumber]; ok {
			return accumulator
		} else {
			usedLines[lineNumber] = 1
		}
		lineSplit := strings.Split(lines[lineNumber], " ")
		operation := lineSplit[0]
		amount, _ := strconv.Atoi(lineSplit[1])

		if operation == "nop" {
			lineNumber = lineNumber+1
			continue
		}
		if operation == "acc" {
			accumulator = accumulator + amount
			lineNumber = lineNumber+1
			continue
		}
		if operation == "jmp" {
			lineNumber = lineNumber + amount
			continue
		}
	}
	return 0
}

func main()  {
	fmt.Println(getAccumulator("input.txt"))
}
