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

	var stringArray []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringArray = append(stringArray, scanner.Text())
	}
	return stringArray
}

func sortString(sl []string) []string {
	var rt []string
	for _, s := range sl {
		cl := strings.Split(s, "")
		sort.Strings(cl)
		rt = append(rt, strings.Join(cl,""))
	}
	return rt
}

func solveProblem(input string) int {
	var count int
	lines := readLines(input)
	numbers := map[string]string{
		"abcdefg": "8",
		"bcdef": "5",
		"acdfg": "2",
		"abcdf": "3",
		"abd": "7",
		"abcdef": "9",
		"bcdefg": "6",
		"abef": "4",
		"abcdeg": "0",
		"ab": "1",
	}

	for _, line := range lines {
		var countString string
		lineArray := strings.Split(line, " | ")
		afterPart := lineArray[1]
		afSplit := strings.Split(afterPart, " ")
		afSplit = sortString(afSplit)
		for _, a := range afSplit {
			if _, ok := numbers[a]; ok {
				fmt.Println(a)
				fmt.Println(numbers[a])
				countString = countString + numbers[a]
			}
		}
		countInt,_ := strconv.Atoi(countString)
		count = count + countInt
	}

	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}