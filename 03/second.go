package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines2(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var stringArray [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var chars []string
		for _, x := range scanner.Text(){
			chars = append(chars, string(x))
		}
		stringArray = append(stringArray, chars)
	}

	return stringArray
}


func getTrees(){

}


func main() {
	paths := readLines2("input.txt")

	var allNumbers []int

	treePath := []int{
		1,
		3,
		5,
		7,
	}
	var increment int
	for _, k := range treePath {
		if k == 12 {
			increment = 1
		} else {
			increment = k
		}
		tmpNum := increment + 1
		trees := 0
		lineLenght := len(paths[0])
		var lineNumber int
		for n, _ := range paths {
			lineNumber = n
			if n >= 1 {
				if string(paths[lineNumber][tmpNum-1]) == "#" {
					trees++
				}
				tmpNum = tmpNum + increment
				if tmpNum > lineLenght {
					tmpNum = tmpNum - lineLenght
				}
			}
		}
		allNumbers = append(allNumbers, trees)
		fmt.Println(trees)
	}

	increment = 1
	lineNumber := 0
	var tree int
	for n, _ := range paths {
		if n >= 2 {
			if increment > len(paths[lineNumber]){
				increment = increment - len(paths[lineNumber])
				lineNumber++
			}
			if paths[lineNumber][increment-1] == "#" {
				tree++
			}
			increment++
			lineNumber++
			lineNumber++
			if lineNumber >= len(paths){
				break
			}
		}

	}
	fmt.Println(tree)
	allNumbers = append(allNumbers, tree)

	asd := 1
	for _, a := range allNumbers {
		asd = asd * a
	}
	fmt.Println(asd)
}

