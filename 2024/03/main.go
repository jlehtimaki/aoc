package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Pair struct {
	x int
	y int
}

func reads1(path string) []Pair {
	var pairArray []Pair
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// need to fetch all integers inside mul()
		regex := regexp.MustCompile(`mul\(\b[0-9]+,\b[0-9]+\)`)
		result := regex.FindAllString(scanner.Text(), -1)

		for _, x := range result {
			regex := regexp.MustCompile(`\d+`)
			result := regex.FindAllString(x, -1)
			xInt, _ := strconv.Atoi(result[0])
			yInt, _ := strconv.Atoi(result[1])
			pairArray = append(pairArray, Pair{x: xInt, y: yInt})
		}
	}
	return pairArray
}

func reads2(path string) []string {
	var input string
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text()
	}

	// check all mul(), don't and do()
	regex := regexp.MustCompile(`mul\(\b[0-9]+,\b[0-9]+\)|don't\(\)|do\(\)`)
	muls := regex.FindAllString(input, -1)
	return muls
}

func s1(input string) int {
	var count int
	pairArray := reads1(input)
	for _, pair := range pairArray {
		count += pair.x * pair.y
	}
	return count
}

func s2(input string) int {
	var count int
	enabled := true
	output := reads2(input)

	for _, x := range output {
		switch x {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				count += getValue(x)
			}
		}
	}

	return count
}

func getValue(mul string) int {
	regex := regexp.MustCompile(`\d+`)
	result := regex.FindAllString(mul, -1)
	xInt, _ := strconv.Atoi(result[0])
	yInt, _ := strconv.Atoi(result[1])

	return xInt * yInt
}

func main() {
	fmt.Printf("P1: %d \n", s1("input.txt"))
	fmt.Printf("P2: %d \n", s2("input.txt"))
}
