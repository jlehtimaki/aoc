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
	var e []string
	var g []string
	var bitArray []int
	for y, line := range lines {
		lineSplit := strings.Split(line, "")
		for n, l := range lineSplit {
			if y == 0 {
				bitArray = append(bitArray, 0)
				g = append(g, "0")
				e = append(e, "0")
			}
			if l == "1" {
				bitArray[n] = bitArray[n] + 1
			}
			if bitArray[n] > (len(lines) / 2){
				g[n] = "1"
				e[n] = "0"
			} else {
				g[n] = "0"
				e[n] = "1"
			}
		}
	}
	eString := strings.Join(e, "")
	gString := strings.Join(g, "")
	fmt.Println(eString)
	fmt.Println(gString)
	eInt,_ := strconv.ParseInt(eString,2, 64)
	gInt,_ := strconv.ParseInt(gString, 2, 64)
	return int(eInt * gInt)
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}