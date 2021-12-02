package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

)

var (
	lines = readLines("input.txt")
	dpMap = map[int]int{}
)

func readLines(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var intArray []int
	intArray = append(intArray, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineInt,_ := strconv.Atoi(scanner.Text())
		intArray = append(intArray, lineInt)
	}
	sort.Ints(intArray)
	increasedLastNum := intArray[len(intArray)-1]+3
	intArray = append(intArray, increasedLastNum)
	return intArray
}

func dp(i int) int{
	if i == len(lines) - 1 {
		return 1
	}
	if _, ok := dpMap[i]; ok {
		return dpMap[i]
	}
	ans := 0
	for x := i+1 ; x < len(lines) ; x++ {
		if lines[x]-lines[i]<=3 {
			ans = ans + dp(x)
		}
	}
	dpMap[i] = ans
	return ans
}

func main()  {
	i := 0
	fmt.Println(dp(i))
}
