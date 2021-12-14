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

	var stringArray []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringArray = append(stringArray, scanner.Text())
	}
	return stringArray
}

func flashes(octos [10][10]int) ([10][10]int, int){
	var flash int
	for true {
		for yn, y := range octos {
			for xn, x := range y {
				if x == -1 || x > 9{
					//fmt.Printf("%d,%d\n",yn,xn)
					octos[yn][xn] = 0
					if yn == 0 {
						if xn == 0 {
							if octos[yn+1][xn] != 0 {
								octos[yn+1][xn] = octos[yn+1][xn] + 1
							}
							for i:=0; i <= 1; i++{
								if octos[yn+i][xn+1] != 0 {
									octos[yn+i][xn+1] = octos[yn+i][xn+1] + 1
								}
							}

						} else if xn == 9 {
							for i:=0; i <= 1 ; i++ {
								if octos[yn+1][xn-i] != 0 {
									octos[yn+1][xn-i] = octos[yn+1][xn-i] + 1
								}
							}
							if octos[yn][xn-1] != 0 {
								octos[yn][xn-1] = octos[yn][xn-1] + 1
							}
						} else {
							for i:=-1; i <= 1; i++ {
								if octos[yn+1][xn+i] != 0 {
									octos[yn+1][xn+i] = octos[yn+1][xn+i] + 1
								}
							}
							if octos[yn][xn-1] != 0{
								octos[yn][xn-1] = octos[yn][xn-1] + 1
							}
							if octos[yn][xn+1] != 0{
								octos[yn][xn+1] = octos[yn][xn+1] + 1
							}
						}
					} else if yn == 9 {
						if xn == 0 {
							//for i:=0; i<=1; i++{
							//
							//}
							for i:=0; i<=1; i++{
								if octos[yn-1][xn+i] != 0 {
									octos[yn-1][xn+i] = octos[yn-1][xn+i] + 1
								}
							}
							if octos[yn][xn+1] != 0 {
								octos[yn][xn+1] = octos[yn][xn+1] + 1
							}
						} else if xn == 9 {
							for i:=0; i<=1; i++{
								if octos[yn-1][xn-i] != 0 {
									octos[yn-1][xn-i] = octos[yn-1][xn-i] + 1
								}
							}
							if octos[yn][xn-1] != 0 {
								octos[yn][xn-1] = octos[yn][xn-1] + 1
							}
						} else {
							for i:=-1; i<=1; i++{
								if octos[yn-1][xn+i] != 0 {
									octos[yn-1][xn+i] = octos[yn-1][xn+i] + 1
								}
							}
							if octos[yn][xn+1] != 0 {
								octos[yn][xn+1] = octos[yn][xn+1] + 1
							}
							if octos[yn][xn-1] != 0 {
								octos[yn][xn-1] = octos[yn][xn-1] + 1
							}
						}
					} else {
						if xn == 0 {
							for i:=-1; i<=1;i++ {
								if octos[yn+i][xn+1] != 0 {
									octos[yn+i][xn+1] = octos[yn+i][xn+1] + 1
								}
							}
							if octos[yn+1][xn] != 0 {
								octos[yn+1][xn] = octos[yn+1][xn] + 1
							}
							if octos[yn-1][xn] != 0 {
								octos[yn-1][xn] = octos[yn-1][xn] + 1
							}
						} else if xn == 9 {
							for i:=-1; i<=1;i++ {
								if octos[yn+i][xn-1] != 0 {
									octos[yn+i][xn-1] = octos[yn+i][xn-1] + 1
								}
							}
							if octos[yn+1][xn] != 0 {
								octos[yn+1][xn] = octos[yn+1][xn] + 1
							}
							if octos[yn-1][xn] != 0 {
								octos[yn-1][xn] = octos[yn-1][xn] + 1
							}
						} else {
							for i:=-1; i<=1;i++ {
								if octos[yn+i][xn-1] != 0 {
									octos[yn+i][xn-1] = octos[yn+i][xn-1] + 1
								}
							}
							for i:=-1; i<=1;i++ {
								if octos[yn+i][xn+1] != 0 {
									octos[yn+i][xn+1] = octos[yn+i][xn+1] + 1
								}
							}
							if octos[yn+1][xn] != 0 {
								octos[yn+1][xn] = octos[yn+1][xn] + 1
							}
							if octos[yn-1][xn] != 0 {
								octos[yn-1][xn] = octos[yn-1][xn] + 1
							}
						}
					}
					flash = flash + 1
				}
			}
		}
		if !contains(octos) {
			break
		}
	}
	return octos, flash
}

func contains(octos [10][10]int) bool {
	for _, y := range octos {
		for _, x := range y {
			if x == -1 || x > 9 {
				return true
			}
		}
	}
	return false
}

func allFlashes(octos [10][10]int) bool {
	for _, y := range octos {
		for _, x := range y {
			if x != 0 {
				return false
			}
		}
	}
	return true
}

func solveProblem(input string) int {
	var count int
	lines := readLines(input)
	var octos [10][10]int
	for n, l := range lines {
		for nn, c := range l {
			number,_ := strconv.Atoi(string(c))
			octos[n][nn] = number
		}
	}

	var tmpInt int
	for i:=0; i < 1000000; i++ {
		for yn, y := range octos {
			for xn, _ := range y {
				octos[yn][xn] = octos[yn][xn] + 1
			}
		}
		octos, tmpInt = flashes(octos)
		for _, y := range octos {
			for _, x := range y {
				fmt.Printf("%d ", x)
			}
			fmt.Printf("\n")
		}
		if allFlashes(octos) {
			return i+1
		}
		fmt.Printf("\n")
		count = count + tmpInt
	}

	for _, y := range octos {
		for _, x := range y {
			fmt.Printf("%d ", x)
		}
		fmt.Printf("\n")
	}
	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}