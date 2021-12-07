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

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func getFuelCost(x int, y int) int {
	var final int
	var min int
	var max int
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	count := 1
	for z:=min; z < max; z++ {
		final = final + count
		count = count + 1
	}

	return final
}

func solveProblem(input string) int {
	var numbers []int
	cache := map[int]int{}
	//cache2 := map[string]int{}
	lines := readLines(input)
	numberString := strings.Split(lines[0],",")
	for _, n := range numberString {
		number, _ := strconv.Atoi(n)
		numbers = append(numbers, number)
	}

	min, max := MinMax(numbers)

	for x:=min; x<=max; x++ {
		var count int
		for _, number := range numbers {
			if _, ok := cache[x]; ok {
				count = count - cache[x]
				continue
			}

			count = count + getFuelCost(x, number)
		}
		cache[x] = count
	}

	fmt.Println(cache)
	var count int
	for _,v := range cache {
		if count == 0 {
			count = v
		}
		if count > v {
			count = v
		}
	}
	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}