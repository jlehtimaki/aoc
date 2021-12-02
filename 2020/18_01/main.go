package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var array []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}
	return array
}

func calcParanth(p string) int {
	var sum int
	//fmt.Println(p)
	for _,ms := range p[:]{
		if ms
	}
	return sum
}

func parseLine(line string) int{
	var sum int
	re := regexp.MustCompile(`\(`)
	if re.MatchString(line){
		sum = calcParanth(line)
	}
	//for i:=0; i < (len(line[:])-4); i++{
	//	firstNum,_ := strconv.Atoi(string(line[i]))
	//}

	return sum
}

func solveProblem(input string) int {
	var sum int
	lines := readLines(input)

	for _, l := range lines{
		p := strings.Trim(l, " ")
		p = strings.ReplaceAll(p,"(","( ")
		p = strings.ReplaceAll(p,")", " )")
		sum = sum + parseLine(p)
	}

	return sum
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}