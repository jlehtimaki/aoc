package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readLines(path string) ([]int, []int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var intA, intB []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "   ")
		a, _ := strconv.Atoi(split[0])
		b, _ := strconv.Atoi(split[1])
		intA = append(intA, a)
		intB = append(intB, b)
	}
	return intA, intB
}

func s1(input string) int {
	var count int
	a, b := readLines(input)

	sort.Ints(a)
	sort.Ints(b)

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			count += a[i] - b[i]
			continue
		}
		count += b[i] - a[i]
	}
	return count
}

func s2(input string) int {
	var count int
	aMap := make(map[int]int)
	a, b := readLines(input)

	for _, i := range b {
		aMap[i]++
	}

	for _, x := range a {
		if aMap[x] > 0 {
			count += x * aMap[x]
		}
	}

	return count
}

func main() {
	fmt.Printf("Problem 1: %d \n", s1("input.txt"))
	fmt.Printf("Problem 2: %d \n", s2("input.txt"))
}

