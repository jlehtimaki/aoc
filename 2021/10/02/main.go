package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func solveProblem(input string) int {
	rules := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
	scoring := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}
	var parsedLines []string
	lines := readLines(input)
	for _, line := range lines {
		foo:
		for true {
			var removed bool
			for k,v := range rules {
				dd := fmt.Sprintf("%s%s",k,v )
				if strings.Contains(line, dd){
					line = strings.ReplaceAll(line, dd, "")
					removed = true
				}
			}

			if !removed {
				break foo
			}
		}
		parsedLines = append(parsedLines, line)
	}

	var unfinishedLines []string
	for _, line := range parsedLines {
		var consists bool
		ss:
		for _,v := range rules {
			if strings.Contains(line, v) {
				consists = true
				break ss
			}
		}
		if !consists {
			unfinishedLines = append(unfinishedLines, line)
		}
	}

	var scores []int
	for _, ul := range unfinishedLines {
		cc := 0
		for i:=len(ul)-1; i>=0; i-- {
			cc = 5 * cc + scoring[string(ul[i])]
		}
		scores = append(scores, cc)
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}