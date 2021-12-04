package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BingoCards struct {
	BingoCards [][5][5]map[string]bool
}

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

func (b *BingoCards) parseData(input []string){
	var bingoCard [5][5]map[string]bool
	var count int
	for n, line := range input {
		if n >= 2 {
			if strings.Contains(line, " "){
				re := regexp.MustCompile("[0-9]+")
				numbers := re.FindAllString(line, -1)
				for nn, number := range numbers {
					tmp := map[string]bool{}
					tmp[number] = false
					bingoCard[count][nn] = tmp
				}
				count = count + 1
			} else {
				b.BingoCards = append(b.BingoCards, bingoCard)
				bingoCard = [5][5]map[string]bool{}
				count = 0
			}
		}
	}
	b.BingoCards = append(b.BingoCards, bingoCard)
}

func (b *BingoCards) checkNumber(number string) {
	for n, bingocard := range b.BingoCards {
		for yn, y := range bingocard {
			for xn, x := range y {
				for k, _ := range x {
					if _, ok := x[number]; ok {
						b.BingoCards[n][yn][xn][k] = true
					}
				}
			}
		}
	}
}

func (b *BingoCards) checkBingo(num string) int{
	numInt, _ := strconv.Atoi(num)
	//Check horizontal
	for _, bingocard := range b.BingoCards {
		for _, y := range bingocard {
			foo := fmt.Sprintf("%v", y)
			if !strings.Contains(foo, "false"){
				return countBingo(bingocard) * numInt
			}
		}
	}

	// check vertical
	for _, bingocard := range b.BingoCards {
		var list []bool
		for x:= 0; x < 5; x++ {
			for y:= 0; y < 5; y++ {
				for _,v := range bingocard[y][x]{
					list = append(list, v)
				}
			}
			stringList := fmt.Sprintf("%v", list)
			if !strings.Contains(stringList, "false"){
				return countBingo(bingocard) * numInt
			}
		}
	}
	return 0
}

func countBingo(card [5][5]map[string]bool) int{
	var count int
	for _, y := range card {
		for _, x := range y {
			for k, v := range x {
				if v == false {
					kInt, _ := strconv.Atoi(k)
					count = count + kInt
				}
			}
		}
	}
	return count
}

func (b *BingoCards) solveProblem(input string) int {
	lines := readLines(input)
	b.parseData(lines)
	numbers := strings.Split(lines[0], ",")
	for n, number := range numbers {
		b.checkNumber(number)
		if n >= 4 {
			v := b.checkBingo(number)
			if v != 0 {
				return v
			}
		}
	}
	return 0
}

func main()  {
	b := BingoCards{}
	fmt.Println(b.solveProblem("input.txt"))
}