package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}


func solveProblem(input string) int {
	lineMap := map[string]int{}
	lines := readLines(input)
	for _, line := range lines {
		lineSplit := strings.Split(line, "->")
		start := strings.TrimSpace(lineSplit[0])
		end := strings.TrimSpace(lineSplit[1])
		x1t,_ := strconv.Atoi(strings.Split(start,",")[0])
		y1t,_ := strconv.Atoi(strings.Split(start,",")[1])
		x2t,_ := strconv.Atoi(strings.Split(end,",")[0])
		y2t,_ := strconv.Atoi(strings.Split(end,",")[1])
		x1 := Min(x1t, x2t)
		x2 := Max(x1t, x2t)
		y1 := Min(y1t, y2t)
		y2 := Max(y1t, y2t)

		//fmt.Println(x1,y1,x2,y2)
		if x1 == x2 || y1 == y2 {
			//fmt.Println(x1,y1,x2,y2)
			for x := x1; x <= x2; x++ {
				for y:=y1; y <= y2; y++ {
					xy := fmt.Sprintf("%d,%d", x, y)
					if _, ok := lineMap[xy]; !ok {
						lineMap[xy] = 0
					}
					lineMap[xy] = lineMap[xy] + 1
				}
			}
		}

	}

	var count int
	for k, v := range lineMap {
		//fmt.Println(lineMap)
		if v > 1 {
			fmt.Printf("%s:%d \n",k,v)
			count = count + 1
		}
	}
	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}