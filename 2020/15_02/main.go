package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	cache = map[int]int{}
)
func newNumber(numbers []int, tn int) int {
	var number int
	tn = tn -1
	lasNum := numbers[tn]
	if _, ok := cache[lasNum]; ok {
		if tn == cache[lasNum] {
			fmt.Println("exiting 1")
			return 0
		}
		number = (tn) - cache[lasNum]
	} else {
		return 0
	}

	return number
}

func solveProblem(input string) int {
	numbersAsString := strings.Split(input,",")
	var tn int
	turns := 30000000
	numbers := []int{}

	// init
	for n, na := range numbersAsString {
		nu,_ := strconv.Atoi(na)
		numbers = append(numbers, nu)
		cache[nu] = n
	}
	tn = len(numbers)

	for tn < turns {
		fmt.Printf("%d\n", tn+1)
		num := newNumber(numbers, tn)
		numbers = append(numbers, num)
		cache[numbers[tn-1]] = tn-1
		tn++
	}
	//fmt.Println(numbers)
	return numbers[turns-1]
}

func main()  {
	fmt.Println(solveProblem("2,0,6,12,1,3"))
}
