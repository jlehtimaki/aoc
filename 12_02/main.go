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
		"R": 1000,
		"L": 1000,
	}
	cDir := 90
	cDir2 := 0
	dirInt := 10
	dirInt2 := 1

	for _, i := range instructions{
		dr := i[0:1]
		am, _ := strconv.Atoi(i[1:])

		fmt.Printf("%s - %d\n", dr, am)
		fmt.Printf("%d: %d - %d: %d\n", cDir, dirInt, cDir2, dirInt2)
		fmt.Println(dirMap)

		if dr == "F"{
			dirMap[cDir] = dirMap[cDir] + (am * dirInt)
			dirMap[cDir2] = dirMap[cDir2] + (am * dirInt2)
			continue
		}

		if dr == "R" {
			if (cDir + am) >= 360 {
				cDir = (cDir + am) - 360
			} else {
				cDir = cDir + am
			}

			if (cDir - 90) < 0 {
				cDir2 = 360 - 90
			} else {
				cDir2 = cDir - 90
			}
			continue
		}

		if dr == "L" {
			if (cDir - am) < 0 {
				cDir = 360 - (am - cDir)
			} else {
				cDir = cDir - am
			}

			if (cDir + 90) >= 360 {
				cDir2 = 0
			} else {
				cDir2 = cDir + 90
			}
			continue
		}

		if degreeMap[dr] == cDir {
			dirInt = dirInt + am
			continue
		}
		if degreeMap[dr] == cDir2 {
			dirInt2 = dirInt2 + am
			continue
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
