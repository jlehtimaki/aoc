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
	var count int
	intArray := [][]int{}
	lines := readLines(input)
	for n, line := range lines {
		intArray = append(intArray, []int{})
		l := strings.Split(line, "")
		for _, ll := range l {
			tInt, _ := strconv.Atoi(ll)
			intArray[n] = append(intArray[n], tInt)
		}
	}

	for yn, y := range intArray {
		for xn, x := range y {
			var adjacent []int
			dummy := true
			if xn == 0 {
				adjacent = append(adjacent, intArray[yn][xn+1])
			} else if xn == (len(y)-1) {
				adjacent = append(adjacent, intArray[yn][xn-1])
			} else {
				adjacent = append(adjacent, intArray[yn][xn-1])
				adjacent = append(adjacent, intArray[yn][xn+1])
			}
			if yn == 0 {
				adjacent = append(adjacent, intArray[yn+1][xn])
			} else if yn == (len(intArray)-1) {
				adjacent = append(adjacent, intArray[yn-1][xn])
			} else {
				adjacent = append(adjacent, intArray[yn-1][xn])
				adjacent = append(adjacent, intArray[yn+1][xn])
			}
			for _, a := range adjacent {
				if x >= a {
					dummy = false
					break
				}
			}
			if dummy {
				count = count + x + 1
			}
		}
	}
	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}