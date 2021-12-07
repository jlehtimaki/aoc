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
	var numbers []int
	lines := readLines(input)
	numbersString := strings.Split(lines[0],",")
	for _, numberString := range numbersString {
		number,_ := strconv.Atoi(numberString)
		numbers = append(numbers, number)
	}

	for x:=0 ; x < 80; x++ {
		for n, number := range numbers {
			if number == 0 {
				number = 8
				numbers = append(numbers, 8)
				numbers[n] = 6
				continue
			}
			number = number - 1
			numbers[n] = number
		}
	}

	return len(numbers)
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}