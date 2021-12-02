package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func readLines2(path string) int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var validPasswords int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if parseLine2(scanner.Text()) {
			validPasswords++
		}
	}
	return validPasswords
}

func createStringArray(s string) []string {
	var tmpString []string
	for _, c := range s {
		tmpString = append(tmpString, string(c))
	}
	return tmpString
}

func parseLine2(line string) (bool) {
	stringArray := strings.Split(line, " ")
	firstNumber, _ := strconv.Atoi(strings.Split(stringArray[0], "-")[0])
	secondNumber, _ := strconv.Atoi(strings.Split(stringArray[0], "-")[1])
	character := strings.ReplaceAll(stringArray[1], ":", "")
	password := createStringArray(stringArray[2])
	firstNumber--
	secondNumber--

	if ( password[firstNumber] == character && password[secondNumber] != character ) || ( password[firstNumber] != character && password[secondNumber] == character ) {
		return true
	}
	return false
}

func main()  {
	fmt.Println(readLines2("input.txt"))
}