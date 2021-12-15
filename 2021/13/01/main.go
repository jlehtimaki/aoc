package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func solveProblem(input string) int {
	var count int
	lines := readLines(input)
	manual := [][]string{}
	coordinates := map[string]int{}
	var xlen int
	var ylen int
	var foldInstructions []map[string]int

	// Get data
	for _, line := range lines {
		if strings.Contains(line, ","){
			coordinates[line] = 0
			lineSplit := strings.Split(line, ",")
			newX, _ := strconv.Atoi(lineSplit[0])
			newY, _ := strconv.Atoi(lineSplit[1])
			if newX > xlen {
				xlen = newX
			}
			if newY > ylen {
				ylen = newY
			}
			continue
		}
		if strings.Contains(line, "fold"){
			re := regexp.MustCompile("[0-9]+")
			number,_ := strconv.Atoi(re.FindAllString(line, -1)[0])
			if strings.Contains(line, "y"){
				foldInstructions = append(foldInstructions, map[string]int{"y": number})
			} else {
				foldInstructions = append(foldInstructions, map[string]int{"x": number})
			}
			continue
		}
	}

	// create manual
	for y:=0; y <= ylen; y++ {
		manual = append(manual, []string{})
		for x:=0; x<= xlen; x++ {
			cord := fmt.Sprintf("%d,%d",x,y)
			if _, ok := coordinates[cord]; ok {
				manual[y] = append(manual[y], "#")
			} else {
				manual[y] = append(manual[y], ".")
			}
		}
	}

	for _, foldInstruction := range foldInstructions {
		for k,v := range foldInstruction {
			if k == "x" {
				tmpManual := [][]string{}
				for y:=0; y < len(manual); y++ {
					for x:=1; x<=v; x++ {
						if manual[y][v+x] == "#"{
							manual[y][v-x] = "#"
						}
					}
					//manual[y] = []string{}
					//tmpSlice := manual[y][:foldX]
					tmpManual = append(tmpManual, manual[y][:v])
				}
				manual = tmpManual
			} else {
				for y:=1; y <= v; y++ {
					for x:=0; x<len(manual[0]); x++ {
						if manual[v+y][x] == "#"{
							manual[v-y][x] = "#"
							continue
						}
					}
				}
				manual = manual[:v]
			}
		}
		break
	}

	// Count seeing dots
	for _, x := range manual{
		line := fmt.Sprintf("%v", x)
		count = count + strings.Count(line, "#")
	}
	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}