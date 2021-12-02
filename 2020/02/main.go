package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var validPasswords int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if parseLine(scanner.Text()) {
			validPasswords++
		}
	}
	return validPasswords
}

func parseLine(line string) (bool) {
	stringArray := strings.Split(line, " ")
	minNumber, _ := strconv.Atoi(strings.Split(stringArray[0], "-")[0])
	maxNumber, _ := strconv.Atoi(strings.Split(stringArray[0], "-")[1])
	character := strings.ReplaceAll(stringArray[1], ":", "")
	password := stringArray[2]

	number := strings.Count(password, character)
	if number >= minNumber && number <= maxNumber {
		return true
	}
	return false
}


func main()  {
	fmt.Println(readLines("input.txt"))
}