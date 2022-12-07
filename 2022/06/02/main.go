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
	line := readLines(input)[0]
	s := ""
	for n, l := range line {
		if len(s) < 14 {
			for true {
				if strings.Contains(s, string(l)) {
					s = s[1:]
					continue
				}
				break
			}
			s += string(l)
			continue
		}
		return n
	}
	return count
}

func main() {
	fmt.Println(solveProblem("input.txt"))
}