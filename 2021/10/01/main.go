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

func solveProblem(input string) int {
	var count int
	rules := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
	scoring := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	var expectedCharacters []string
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

	for _, line := range parsedLines {
		ss:
		for n, c := range line {
			for k,v := range rules {
				if string(c) == v {
					if string(line[n-1]) != k {
						expectedCharacters = append(expectedCharacters, string(c))
						break ss
					}
				}
			}
		}
	}

	for k,v := range scoring{
		c := strings.Count(strings.Join(expectedCharacters," "), k)
		count = count + (c * v)
	}


	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}

//{([(<{}[<>[]}>{[]{[(<()> - Expected ], but found } instead.
//[[<[([]))<([[{}[[()]]] - Expected ], but found ) instead.
//[{[{({}]{}}([{[{{{}}([] - Expected ), but found ] instead.
//[<(<(<(<{}))><([]([]() - Expected >, but found ) instead.
//<{([([[(<>()){}]>(<<{{

//{([(<[}>{{[(
//[[<[())<([[[]]
//[{[{(]}([{[{{}(
//[<(<(<(<))><((
//<{([([[()]>(<<{{


//[({(<()>[[{{<>
//[([])]({[<{<>(
//{([(<[}>{{[(
//(((({}<{<{}{{
//[[<[())<([[[]]
//[{[{(]}([{[{{}(
//{<[]>}<{[{[{{[[
//[<(<(<(<))><((
//<{([([[()]>(<<{{
//<{([{}[<[[[]]]>]




