package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func getAllLetters(lines []string) map[string]int {
	letterMap := map[string]int{}
	re := regexp.MustCompile(`[A-Z]`)
	for _, line := range lines {
		letters := re.FindAllString(line, -1)
		for _, l := range letters {
			if _, ok := letterMap[l]; ok {
				continue
			}
			letterMap[l] = 0
		}
	}
	return letterMap
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func countLine(input []string, lineNumber int) int{
	count := 0
	numbers := input[lineNumber:lineNumber+3]
	for _, n := range numbers {
		nn, _ := strconv.Atoi(n)
		count = count + nn
	}
	return count
}

func solveProblem(input string) int {
	lines := readLines(input)
	var count int
	//var prevValue int
	//prevValue = countLine(lines, 0)
	//lines = remove(lines, 0)
	for true {
		if len(lines) > 3 {
			n1, _ := strconv.Atoi(lines[0])
			n2, _ := strconv.Atoi(lines[3])
			if n2 > n1 {
				count = count + 1
			}
			lines = remove(lines, 0)
		} else {
			break
		}
	}
	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}
