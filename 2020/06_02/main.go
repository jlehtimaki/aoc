package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(path string) []map[string]int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var stringArray []map[string]int
	var tmpArray map[string]int
	var tmpString string
	var count int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			count++
			for _, s := range scanner.Text() {
				tmpString = tmpString+string(s)
			}
			continue
		}
		tmpArray = checkAnswers(tmpString, count)
		stringArray = append(stringArray, tmpArray)
		count = 0
		tmpArray = map[string]int{}
		tmpString = ""
	}
	tmpArray = checkAnswers(tmpString, count)
	stringArray = append(stringArray, tmpArray)
	return stringArray
}

func checkAnswers(fakeString string, count int) map[string]int {
	answer := map[string]int{}
	for _, s := range fakeString[0:len(fakeString)] {
		if strings.Count(fakeString, string(s)) == count {
			answer[string(s)] = 0
		}
	}
	return answer
}


func countRightAnswers(input string) int{
	groups := readLines(input)
	var count int
	for _, k := range groups {
		count = count + len(k)
	}
	return count
}

func main()  {
	fmt.Println(countRightAnswers("input.txt"))
}
