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

	var stringAr []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringAr = append(stringAr, scanner.Text())
	}
	return stringAr
}

func countAdjacent(seats []string) int{
	var count int
	for _, x := range seats {
		for _,s := range x[0:] {
			if string(s) == "#"{
			count++
		}
		}
	}

	return count
}

func getSeatOrder(seatLayout []string) ([]string, int){
	var seatOrder []string
	var count int
	for n, l := range seatLayout {
		seatString := ""
		for sn, s := range l[:] {
			seat := string(s)
			if seat == "." {
				seatString = seatString+"."
				continue
			}

			// get correct seatNumberRanges
			seatMin := 0
			seatMax := 0
			if (sn - 1) < 0 {
				seatMin = sn
			} else {
				seatMin = sn - 1
			}

			if sn + 2 > len(l) {
				seatMax = sn + 1
			} else {
				seatMax = sn + 2
			}

			// get adjacents
			var adjacentCount int
			var tmpArray []string
			if n == 0 {
				tmpArray = append(tmpArray, seatLayout[n][seatMin:seatMax])
				tmpArray = append(tmpArray, seatLayout[n+1][seatMin:seatMax])
				adjacentCount = countAdjacent(tmpArray)
			} else if n >= len(seatLayout) - 1 {
				tmpArray = append(tmpArray, seatLayout[n-1][seatMin:seatMax])
				tmpArray = append(tmpArray, seatLayout[n][seatMin:seatMax])
				adjacentCount = countAdjacent(tmpArray)
			} else {
				tmpArray = append(tmpArray, seatLayout[n-1][seatMin:seatMax])
				tmpArray = append(tmpArray, seatLayout[n][seatMin:seatMax])
				tmpArray = append(tmpArray, seatLayout[n+1][seatMin:seatMax])
				adjacentCount = countAdjacent(tmpArray)
			}

			adjacentCount = countAdjacent(tmpArray)
			if seat == "L" {
				if adjacentCount == 0 {
					seatString=seatString+"#"
					continue
				}
			}

			if seat == "#" {
				if adjacentCount-1 >= 4 {
					seatString=seatString+"L"
					continue
				}
			}
			seatString = seatString+seat
			count++
		}
		seatOrder = append(seatOrder, seatString)
	}
	return seatOrder, count
}

func solveProblem(firstLayout []string) int{
	wantChanged := len(firstLayout) * len(firstLayout[0])
	fmt.Println(wantChanged)
	seatLayout := firstLayout
	var changed int
	for true {
		seatLayout, changed = getSeatOrder(seatLayout)
		if changed == wantChanged {
			fmt.Println(changed)
			break
		} else {
			wantChanged = changed
		}
	}

	fmt.Println("counting lines")
	var count int
	for _, seatLines := range seatLayout {
		for _, seat := range seatLines[0:]{
			if string(seat) == "#" {
				count++
			}
		}
	}

	return count
}

func main()  {
	lines := readLines("input.txt")
	fmt.Println(solveProblem(lines))
}
