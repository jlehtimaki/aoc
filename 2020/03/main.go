package main

import (
	"bufio"
	"fmt"
	"os"
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

func checkForTrees(treeString string) int  {
	trees := 0
	for _, s := range treeString {
		if string(s) == "#" {
			trees++
		}
	}
	return trees
}


func main() {
	lines := readLines("input.txt")
	//increment := []int{1,3,5,7}
	//incrementNumber := 1
	tmpNum := 4
	trees := 0
	lineLenght := len(lines[0])
	for n, line := range lines {
		if n >= 1 {
			if string(line[tmpNum-1]) == "#" {
				trees++
			}
			tmpNum = tmpNum + 3
			if tmpNum > lineLenght {
				tmpNum = tmpNum - lineLenght
			}
 		}
	}

	fmt.Println(trees)
}