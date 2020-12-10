package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readLines(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var array []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineInt,_ := strconv.Atoi(scanner.Text())
		array = append(array, lineInt)
	}
	return array
}

func checkSum(input []int, sum int) (bool, int) {
	i := 0
	count := 0
	for i <= len(input) {
		count = count + input[i]
		if count > sum {
			return false, 0
		}
		if count == sum {
			return true, i
		}
		i++
	}
	return false, 0
}

func solveProblem(input string, sum int) int {
	lines := readLines(input)
	for n, _ := range lines {
		ok, num := checkSum(lines[n:], sum)
		if ok {
			tmpArray := lines[n:num+n]
			sort.Ints(tmpArray)
			return tmpArray[0] + tmpArray[len(tmpArray)-1]
		}
	}

	return 0
}

func main()  {
	fmt.Println(solveProblem("input.txt", 21806024))
}