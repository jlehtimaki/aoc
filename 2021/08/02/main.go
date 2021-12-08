package main

import (
	"bufio"
	"fmt"
	"os"
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
	numbers := map[string]int{
		"cd": 1,
		"bcdf": 4,
		"acf": 7,
		"abcdefg": 8,
	}

	for _, line := range lines {
		lineArray := strings.Split(line, " | ")
		afterPart := lineArray[1]
		afSplit := strings.Split(afterPart, " ")
		for k,_ := range numbers {
			for _, a := range afSplit {
				if len(k) == len(a) {
					count = count + 1
					continue
				}
			}
		}
	}

	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}