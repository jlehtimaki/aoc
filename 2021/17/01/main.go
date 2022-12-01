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

func getRanges(s []string) ([]int, []int) {
	var xRange []int
	var yRange []int

	xRangeString := strings.Split(strings.ReplaceAll(s[2],"x=",""),"..")
	yRangeString := strings.Split(strings.ReplaceAll(s[3],"y=",""),"..")

	for _, x := range xRangeString {
		xInt, _ := strconv.Atoi(x)
		xRange = append(xRange, xInt)
	}

	for _, y := range yRangeString {
		yInt, _ := strconv.Atoi(y)
		yRange = append(yRange, yInt)
	}

	return xRange, yRange
}

func shootProbe(x, y int, targetArea map[string]int, maxX, maxY int) (bool, int) {
	var highestPoint int
	xx := 0
	yy := 0
	for true {
		xx += x
		yy += y
		y -= 1

		if x > 0 {
			x -= 1
		}
		coord := fmt.Sprintf("%d,%d",xx,yy)
		if yy > highestPoint {
			highestPoint = yy
		}
		if _, ok := targetArea[coord]; ok {
			return true, highestPoint
		}

		//time.Sleep(2*time.Second)
		if yy < maxY || xx > maxX {
			return false, highestPoint
		}
	}
	return false, highestPoint
}

func solveProblem(input string) int {
	var xRange []int
	var yRange []int
	targetArea := map[string]int{}
	lines := readLines(input)
	lineSplit := strings.Split(strings.ReplaceAll(lines[0],",","")," ")


	xRange, yRange = getRanges(lineSplit)


	for y:=yRange[0]; y<=yRange[1]; y++ {
		for x:=xRange[0]; x<=xRange[1]; x++ {
			coordinate := fmt.Sprintf("%d,%d", x,y)
			targetArea[coordinate] = 0
		}
	}

	//fmt.Println(shootProbe(6,9,targetArea, xRange[1],yRange[1]))

	trueList := map[int]int{}
	var y int
	var highestPoint int
	for true {
		for x:=0; x <= xRange[1]; x++ {
			condition, newHighestPoint := shootProbe(x,y,targetArea,xRange[1], yRange[1])
			if  condition {
				if highestPoint < newHighestPoint {
					highestPoint = newHighestPoint
				}
				trueList[y] = 0
			}
		}
		if y == 2000 {
			break
		}
		y += 1
	}

	return highestPoint
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}