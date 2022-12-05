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
		if strings.Contains(scanner.Text(), "move") {
			stringArray = append(stringArray, scanner.Text())
		}
	}
	return stringArray
}

// func removeFromSlice(slice []string, s int) []string {
// 	fmt.Println(len(slice))
// 	if len(slice) == 1 {
// 		return []string{}
// 	}
// 	return append(slice[:s], slice[s+1:]...)
// }

func solveProblem(input string, stacks [][]string) string {
	var stack string
	lines := readLines(input)
	for _, l := range lines {
		lSplits := strings.Split(string(l), " ")
		amount, _ := strconv.Atoi(lSplits[1])
		from, _ := strconv.Atoi(lSplits[3])
		to, _ := strconv.Atoi(lSplits[5])
		from--
		to--
		stacks[to] = append(stacks[from][:amount:amount], stacks[to]...)
		stacks[from] = stacks[from][amount:]
	}
	for _, l := range stacks {
		stack += l[0]
	}
	return stack
}

func main() {
	stacks := [][]string{
		{"N", "V", "C", "S"},
		{"S", "N", "H", "J", "M", "Z"},
		{"D", "N", "J", "G", "T", "C", "M"},
		{"M", "R", "W", "J", "F", "D", "T"},
		{"H", "F", "P"},
		{"J", "H", "Z", "T", "C"},
		{"Z", "L", "S", "F", "Q", "R", "P", "D"},
		{"W", "P", "F", "D", "H", "L", "S", "C"},
		{"Z", "G", "N", "F", "P", "M", "S", "D"},
	}
	fmt.Println(solveProblem("input.txt", stacks))
}