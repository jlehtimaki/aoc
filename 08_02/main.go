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

func getAccumulator(input string) int{
	oldLines := readLines(input)
	//fmt.Println(oldLines)


	for n, _ := range oldLines{
		// vars
		t := 0
		lineNumber := 0
		accumulator := 0
		lines := make([]string, len(oldLines))
		copy(lines, oldLines)

		// splitting one line
		tmpLineSplit := strings.Split(lines[n], " ")

		if tmpLineSplit[0] == "nop" {
			lines[n] = fmt.Sprintf("jmp %s",tmpLineSplit[1])
		} else if tmpLineSplit[0] == "jmp"{
			lines[n] = fmt.Sprintf("nop %s",tmpLineSplit[1])
		} else {
			continue
		}
		for t < 1000 {
			t++
			if lineNumber == len(lines){
				return accumulator
			}
			lineSplit := strings.Split(lines[lineNumber], " ")
			operation := lineSplit[0]
			amount, _ := strconv.Atoi(lineSplit[1])

			if operation == "nop" {
				lineNumber = lineNumber+1
				continue
			}
			if operation == "acc" {
				accumulator = accumulator + amount
				lineNumber = lineNumber+1
				continue
			}
			if operation == "jmp" {
				lineNumber = lineNumber + amount
				continue
			}
		}
	}
	return 0
}

func main()  {
	fmt.Println(getAccumulator("input.txt"))
}
