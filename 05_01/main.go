package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func getId(minNum int, maxNum int, characters string) int {
	var num int
	for _, char := range characters {
		num = (minNum + maxNum) / 2
		upper, _ := regexp.MatchString(`B|R`, string(char))
		if upper{
			num++
			minNum = num
		} else {
			maxNum = num
		}
	}
	return num
}

func getSeatId(id string) int{
	num := 8

	// Get row
	minNum := 0
	maxNum := 127
	rowCharacters := id[0:7]
	num = num * getId(minNum, maxNum, rowCharacters)

	// Get column
	rowCharacters = id[7:10]
	minNum = 0
	maxNum = 7
	num = num + getId(minNum, maxNum, rowCharacters)


	return num
}


func main()  {
	lines := readLines("input.txt")
	biggest := 0
	for _, line := range lines {
		gs := getSeatId(line)
		if  gs > biggest {
			biggest = gs
		}
	}
	fmt.Println(biggest)
}
