package main

import (
	"bufio"
	"fmt"
	"os"
)


func readLines(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var stringAr [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var tmpStringAr []string
		for _, s := range scanner.Text(){
			tmpStringAr = append(tmpStringAr, string(s))
		}
		stringAr = append(stringAr, tmpStringAr)
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

func getSeatOrder(seatLayout [][]string) ([]string, bool){
	rLen := len(seatLayout)
	cLen := len(seatLayout[0])
	var seatOrder []string
	var changed bool
	for r, row := range seatLayout {
		for c, _ := range row {
			for dr:=-1; dr<=1; dr++ {
				for dc:=-1; dc<=1; dc++ {
					if dr == 0 && dc == 0 {
						continue
					}
					rr = r+dr
					cc = c+dc

					if 0<=rr<rLen &&
				}
			}
		}
		seatOrder = append(seatOrder, seatString)
	}
	return seatOrder, changed
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
