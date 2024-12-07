package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func readLines(path string) map[int][]int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	intMap := make(map[int][]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strArray := strings.Split(scanner.Text(), ":")
		value := strings.TrimSpace(strArray[0])
		valueInt, _ := strconv.Atoi(value)
		nums := strings.Split(strArray[1], " ")
		var numsInt []int
		for _, num := range nums {
			intNum, _ := strconv.Atoi(num)
			numsInt = append(numsInt, intNum)
		}
		intMap[valueInt] = numsInt
	}
	return intMap
}

func generateOps1(n int) [][]string {
	opsOptions := []string{"+", "*", "||"}
	total := int(math.Pow(2, float64(n))) // 2^(n)
	opsCombinations := make([][]string, 0, total)

	for i := 0; i < total; i++ {
		combo := make([]string, n)
		x := i
		for j := 0; j < n; j++ {
			digit := x % 2
			x = x / 2
			combo[j] = opsOptions[digit]
		}
		opsCombinations = append(opsCombinations, combo)
	}

	return opsCombinations
}

func generateOps2(n int) [][]string {
	opsOptions := []string{"+", "*", "||"}
	total := int(math.Pow(3, float64(n))) // 3^(n)
	opsCombinations := make([][]string, 0, total)

	for i := 0; i < total; i++ {
		combo := make([]string, n)
		x := i
		for j := 0; j < n; j++ {
			digit := x % 3
			x = x / 3
			combo[j] = opsOptions[digit]
		}
		opsCombinations = append(opsCombinations, combo)
	}

	return opsCombinations
}

func findMatch(value int, nums []int, problem string) bool {
	var ops [][]string
	if problem == "p1" {
		ops = generateOps1(len(nums) - 1)
		for _, op := range ops {
			result := nums[0]
			for i, num := range nums[1:] {
				if op[i] == "*" {
					result *= num
					continue
				}
				result += num
			}
			if result == value {
				return true
			}
		}
	} else {
		ops = generateOps2(len(nums) - 1)
		for _, op := range ops {
			result := nums[0]
			for i, num := range nums[1:] {
				switch op[i] {
				case "+":
					result += num
				case "*":
					result *= num
				case "||":
					digitString, err := strconv.Atoi(fmt.Sprintf("%d%d", result, num))
					if err != nil {
						fmt.Println("Error converting to int")
					}
					result = digitString
				}
			}
			if result == value {
				return true
			}
		}
	}

	return false
}

func s1(input string) int {
	var count int
	lines := readLines(input)

	for value, nums := range lines {
		if findMatch(value, nums, "p1") {
			count += value
		}
	}

	return count
}

func s2(input string) int {
	var count int
	lines := readLines(input)

	for value, nums := range lines {
		wg.Add(1)
		go func(value int, nums []int) {
			defer wg.Done()
			// Use the copy of g and do not modify the original g
			if findMatch(value, nums, "p2") {
				count += value
			}
		}(value, nums)
	}

	wg.Wait()
	return count
}

func main() {
	fmt.Printf("P1: %d \n", s1("input.txt"))
	fmt.Printf("P2: %d \n", s2("input.txt"))
}
