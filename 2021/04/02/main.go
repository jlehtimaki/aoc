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
	Numbers []string
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

func (b *BingoCards) checkBingo(num int) int{
	numInt, _ := strconv.Atoi(b.Numbers[num])
	//Check horizontal
	//if len(b.BingoCards) <= 2 {
	//	printer(b.BingoCards)
	//}
	// check if only one left count the score
	//fmt.Println(len(b.BingoCards))
	if len(b.BingoCards) == 1 {
		return countBingo(b.BingoCards[0]) * numInt
	}

	foo:
	for n, bingocard := range b.BingoCards {
		fmt.Println(len(bingocard))
		for _, bingoRow := range bingocard {
			bRowString := fmt.Sprintf("%v", bingoRow)
			if !strings.Contains(bRowString, "false"){
				b.BingoCards = remove(b.BingoCards, n)
				break foo
			}
		}

		for x:=0; x<=len(bingocard); x++ {
			var bingoColumn string
			for y:=0; y<=len(bingocard); y++ {
				var number string
				for k,_ := range bingocard[x][y] {
					number = k
				}
				bingoColumn = fmt.Sprintf("%s %s ", bingoColumn, number)
			}
			if !strings.Contains(bingoColumn, "false"){
				b.BingoCards = remove(b.BingoCards, n)
				break foo
			}
		}
	}


	//fmt.Println(len(b.BingoCards))
	//printer(b.BingoCards)
	return 0
}

func printer(dodo [][5][5]map[string]bool) {
	for _, d := range dodo {
		for _, y := range d {
			for _, x := range y {
				fmt.Print(x)
			}
			fmt.Printf("\n")
		}
		fmt.Printf("-------------------\n")
	}
}


func remove(s [][5][5]map[string]bool, i int) [][5][5]map[string]bool {
	return append(s[:i], s[i+1:]...)
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
	b.Numbers = strings.Split(lines[0], ",")
	for n, number := range b.Numbers {
		b.checkNumber(number)
		if n >= 4 {
			v := b.checkBingo(n)
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