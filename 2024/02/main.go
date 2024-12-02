package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readLines(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var stringArray [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var sI []int
		sA := strings.Split(scanner.Text(), " ")
		for _, s := range sA {
			i, _ := strconv.Atoi(s)
			sI = append(sI, i)
		}
		stringArray = append(stringArray, sI)
	}
	return stringArray
}

func s1(input string) int {
	var count int
	lines := readLines(input)
	for _, line := range lines {
		if isGood(line) {
			count++
		}
	}
	return count
}

func s2(input string) int {
	var count int
	lines := readLines(input)
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			tmpLine := remove(line, i)
			if isGood(tmpLine) {
				count++
				break
			}
		}
	}
	return count
}

func isGood(line []int) bool {
	if !sort.IntsAreSorted(line) {
		if !isReverseSorted(line) {
			return false
		}
	}

	for i := 0; i < len(line)-1; i++ {
		diff := math.Abs(float64(line[i]) - float64(line[i+1]))

		if 1 > diff || diff > 3 {
			return false
		}
	}
	return true
}

func remove(s []int, i int) []int {
	newSlice := append([]int(nil), s...)
	return append(newSlice[:i], newSlice[i+1:]...)
}

func isReverseSorted(slice []int) bool {
	return sort.SliceIsSorted(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
}

func main() {
	fmt.Printf("P1: %d \n", s1("input.txt"))
	fmt.Printf("P2: %d \n", s2("input.txt"))
}
