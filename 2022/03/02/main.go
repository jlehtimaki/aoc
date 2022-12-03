package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
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
	alphabet := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}
	lines := readLines(input)
	n := 0
	ar := []string{}
	for _, l := range lines {
		n += 1
		ar = append(ar, l)
		if n == 3 {
			c := findCommonChar(ar)
			n = 0
			ar = []string{}
			if unicode.IsUpper(c) {
				count += alphabet[strings.ToLower(string(c))] + 26
				continue
			}
			count += alphabet[string(c)]
		}
	}
	return count
}

func findCommonChar(a []string) rune {
	for _, y := range a[0] {
		t, _ := regexp.MatchString(string(y), a[1])
		if !t {
			continue
		}
		t2, _ := regexp.MatchString(string(y), a[2])
		if t2 {
			return y
		}
	}
	return rune(0)
}

func main() {
	fmt.Println(solveProblem("input.txt"))
}