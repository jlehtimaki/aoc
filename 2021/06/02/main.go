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
	var numberList []int
	numbers := map[int]int{}
	lines := readLines(input)
	numbersString := strings.Split(lines[0],",")
	for _, numberString := range numbersString {
		number,_ := strconv.Atoi(numberString)
		numberList = append(numberList, number)
	}

	for _, number := range numberList {
		if _, ok := numbers[number]; !ok {
			numbers[number] = 0
		}
		numbers[number] = numbers[number] + 1
	}


	for x:=0 ; x < 256; x++ {
		dummy := map[int]int{}
		for k, v := range numbers{
			if k == 0 {
				dummy[6] = dummy[6] + v
				dummy[8] = dummy[8] + v
			} else {
				dummy[k-1] = dummy[k-1] + v
			}
		}
		numbers = dummy
		fmt.Println(numbers)
	}

	var count int
	for _, v := range numbers {
		count = count + v
	}
	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}