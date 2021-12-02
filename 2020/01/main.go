package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]int, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	tmpInt, _ := strconv.Atoi(scanner.Text())
        lines = append(lines, tmpInt)
    }
    return lines, scanner.Err()
}

func removeIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)

}

func findInt(a []int, x int) (int) {
	for _, n := range a {
		if x == n {
			return n
		}
	}
	return 0
}

func findNumber() (int) {
	wantedNumber := 2020
	lines, _ := readLines("input2.txt")

	for n, l := range lines {
		newList := removeIndex(lines, n)
		for _, x := range newList{
			calc := wantedNumber - ( l + x)

			tmpInt := findInt(newList, calc)
			if tmpInt != 0 {
				fmt.Println(x)
				fmt.Println(l)
				fmt.Println(tmpInt)
				return x*l*tmpInt
			}
		}
	}
	return 0
}

func main()  {
	fmt.Println(findNumber())
}
