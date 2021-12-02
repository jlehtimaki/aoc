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
	lines := readLines(input)
	var prevLine int
	var count int
	for n, l := range lines {
		if n == 0 {
			continue
		}
		lineInt,_ := strconv.Atoi(l)
		if lineInt >= prevLine {
			count = count + 1
		}
		prevLine = lineInt
	}
	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}
