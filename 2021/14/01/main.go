package main

import (
	"bufio"
	"fmt"
	"os"
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

func getPairs(s string) []string {
	var pairs []string

	for n, c := range s {
		if n == 0 {
			continue
		}
		pairs = append(pairs, string(s[n-1])+string(c))
	}

	return pairs
}

func solveProblem(input string) int {
	//var count int
	lines := readLines(input)
	var polymerTemplate string
	var insertionList []map[string]string

	for n, line := range lines {
		if n == 0 {
			polymerTemplate = line
			continue
		}
		if strings.Contains(line, "->") {
			sSplit := strings.Split(line," -> ")
			insertionList = append(insertionList, map[string]string{sSplit[0]: sSplit[1]})
			continue
		}
	}

	for i:=0; i<=9; i++ {
		index := 1
		pairs := getPairs(polymerTemplate)
		for _, pair := range pairs {
			for _, il := range insertionList {
				if _, ok := il[pair]; ok {
					polymerTemplate = polymerTemplate[:index] + il[pair] + polymerTemplate[index:]
					index = index + 2
					break
				}
			}
		}
	}

	letters := map[string]int{}
	var most int
	var least int

	for _, c := range polymerTemplate {
		if _, ok := letters[string(c)]; !ok {
			letters[string(c)] = 1
			continue
		}
		letters[string(c)] = letters[string(c)] + 1
	}

	for _,v := range letters {
		if most == 0  && least == 0{
			most = v
			least = v
		}

		if v > most {
			most = v
		}

		if v < least {
			least = v
		}
	}

	return most - least
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}