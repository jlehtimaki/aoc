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
	lines := readLines(input)
	var x int
	var y int
	var a int
	for _, l := range lines {
		n, _ := strconv.Atoi(strings.Split(l, " ")[1])
		if strings.Contains(l, "forward") {
			x = n + x
			y = y + a * n
			continue
		}
		if strings.Contains(l, "up") {
			a = a - n
			continue
		}
		a = a + n
	}
	return x * y
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}
