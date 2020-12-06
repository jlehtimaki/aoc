package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) []map[string]int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var stringArray []map[string]int
	tmpMap := map[string]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != ""{
			for _, x := range scanner.Text()[0:len(scanner.Text())] {
				tmpMap[string(x)] = 1
			}
		} else {
			stringArray = append(stringArray, tmpMap)
			tmpMap = map[string]int{}
		}
	}
	stringArray = append(stringArray, tmpMap)
	return stringArray
}

func countRightAnswers(input string) int{
	groups := readLines(input)
	var count int
	for _, k := range groups {
		for x, _ := range k {
			count = count + len(x)
		}
	}
	return count
}

func main()  {
	fmt.Println(countRightAnswers("input.txt"))
}
