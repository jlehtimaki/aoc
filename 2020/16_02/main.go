package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var array []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}
	return array
}

func getNumbers(numberRange []string) []int {
	var numbers []int

	for _, numberString := range numberRange {
		if strings.Contains(numberString, "-"){
			nums := strings.Split(numberString, "-")
			fn,_ := strconv.Atoi(nums[0])
			sn,_ := strconv.Atoi(nums[1])
			for i:=fn; i<=sn; i++{
				numbers = append(numbers, i)
			}
		} else {
			sa := strings.Split(numberString, ",")
			for _, n := range sa {
				nn,_ := strconv.Atoi(n)
				numbers = append(numbers, nn)
			}
		}
	}
	sort.Ints(numbers)
	return numbers
}

func parseStrings(lines []string) ([]int, []int, []int) {
	var constraints []int
	var yt []int
	var nt []int

	for n, line := range lines {
		if strings.Contains(line, "your ticket"){
			yt = getNumbers(lines[n+1:n+2])
			continue
		}
		if strings.Contains(line, "nearby tickets"){
			nt = getNumbers(lines[n+1:])
			continue
		}
		if strings.Contains(line, ":"){
			secondColumn := strings.Split(line, ": ")[1]
			numberParse := strings.Split(secondColumn, " or ")
			constraints = append(constraints, getNumbers(numberParse)...)
		}
	}

	return constraints, yt, nt
}

func contains(list []int, x int) bool {
	for _, item := range list {
		if item == x {
			return true
		}
	}
	return false
}

func solveProblem(input string) int {
	lines := readLines(input)

	constraints, yt, nt := parseStrings(lines)

	var falseNumbers []int
	var nums int

	for _, ytt := range yt {
		if contains(constraints, ytt){
			continue
		}
		falseNumbers = append(falseNumbers, ytt)
	}

	for _, ntt := range nt {
		if contains(constraints, ntt){
			continue
		}
		falseNumbers = append(falseNumbers, ntt)
	}

	fmt.Println(falseNumbers)
	for _, fn := range falseNumbers {
		nums = nums + fn
	}
	return nums
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}