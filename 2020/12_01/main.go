package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var stringAr []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringAr = append(stringAr, scanner.Text())
	}
	return stringAr
}

func solveProblem(input string) int {

	instructions := readLines(input)

	dirMap := map[int]int{}
	degreeMap := map[string]int{
		"N": 0,
		"E": 90,
		"S": 180,
		"W": 270,
	}
	cDir := 90

	for _, i := range instructions{
		dr := i[0:1]
		am, _ := strconv.Atoi(i[1:])


		if dr == "F"{
			dirMap[cDir] = dirMap[cDir] + am
			continue
		}

		if dr == "N" || dr == "S" || dr == "E" || dr == "W" {
			dirMap[degreeMap[dr]] = dirMap[degreeMap[dr]] + am
			continue
		}

		if dr == "R" {
			if (cDir + am) >= 360 {
				cDir = (cDir + am) - 360
				continue
			}
			cDir = cDir + am
		}

		if dr == "L" {
			if (cDir - am) < 0 {
				cDir = 360 - (am - cDir)
				continue
			}
			cDir = cDir - am
		}
	}

	var lot int
	fmt.Println(dirMap)
	if dirMap[0] >= dirMap[180] {
		lot = dirMap[0] - dirMap[180]
	} else {
		lot = dirMap[180] - dirMap[0]
	}
	var lat int
	if dirMap[90] >= dirMap[270] {
		lat = dirMap[90] - dirMap[270]
	} else {
		lat =  dirMap[270] - dirMap[90]
	}


	return lat + lot
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}
