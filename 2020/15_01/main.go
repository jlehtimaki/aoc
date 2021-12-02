package main

import (
	"fmt"
	"strconv"
	"strings"
)

func newNumber(numbers []int) int {
	var number int
	turns := []int{}
	lasNum := numbers[len(numbers)-1]
	for tn, n := range numbers {
		if n == lasNum {
			turns = append(turns, tn)
		}
	}

	if len(turns) < 2 {
		return 0
	}

	number = turns[len(turns)-1] - turns[len(turns)-2]
	return number
}

func solveProblem(input string) int {
	numbersAsString := strings.Split(input,",")
	var tn int
	turns := 2020
	numbers := []int{}

	// init
	for _, na := range numbersAsString {
		nu,_ := strconv.Atoi(na)
		numbers = append(numbers, nu)
	}
	tn = len(numbers)

	for tn < turns {
		numbers = append(numbers, newNumber(numbers))
		tn++
	}
	return numbers[turns-1]
}

func main()  {
	fmt.Println(solveProblem("2,0,6,12,1,3"))
}
