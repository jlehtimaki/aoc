package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readLines(path string) (map[int][]int, [][]int) {
	var pageNumbers [][]int
	orderMap := make(map[int][]int)

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		if strings.Contains(scanner.Text(), "|") {
			orderSplit := strings.Split(scanner.Text(), "|")
			o1, _ := strconv.Atoi(orderSplit[0])
			o2, _ := strconv.Atoi(orderSplit[1])
			orderMap[o1] = append(orderMap[o1], o2)
			continue
		}
		if strings.Contains(scanner.Text(), ",") {
			var pageSplitInt []int
			pageSplit := strings.Split(scanner.Text(), ",")
			for _, p := range pageSplit {
				pInt, _ := strconv.Atoi(p)
				pageSplitInt = append(pageSplitInt, pInt)
			}
			pageNumbers = append(pageNumbers, pageSplitInt)
		}
	}
	return orderMap, pageNumbers
}

func checkOrder(on, pn []int) bool {
	for _, n := range pn {
		if !slices.Contains(on, n) {
			return false
		}
	}
	return true
}

func middleNumber(x []int) int {
	y := len(x) / 2
	return x[y]
}

func isGood(pageN []int, orderMap map[int][]int) bool {
	for i := 1; i < len(pageN); i++ {
		if _, ok := orderMap[pageN[i-1]]; !ok {
			return false
		}

		if !checkOrder(orderMap[pageN[i-1]], pageN[i:]) {
			return false
		}
	}
	return true
}

func goodPagesBadPages(orderMap map[int][]int, pageNumbers [][]int) ([][]int, [][]int) {
	var goodPages [][]int
	var badPages [][]int

	for _, pageN := range pageNumbers {
		if isGood(pageN, orderMap) {
			goodPages = append(goodPages, pageN)
			continue
		}
		badPages = append(badPages, pageN)
	}

	return goodPages, badPages
}

func swapElements(bp []int, orderMap map[int][]int) []int {
	// try to do bubble sort
	for i := 0; i < len(bp); i++ {
		if !checkOrder(orderMap[bp[i]], bp[i+1:]) {
			bp[i], bp[i+1] = bp[i+1], bp[i]
			if isGood(bp, orderMap) {
				return bp
			}
		}
	}

	// recursive
	return swapElements(bp, orderMap)
}

func s1(input string) int {
	var count int
	orderMap, pageNumbers := readLines(input)

	goodPages, _ := goodPagesBadPages(orderMap, pageNumbers)
	for _, gp := range goodPages {
		count += middleNumber(gp)
	}
	return count
}

func s2(input string) int {
	var count int
	var goodPages [][]int
	orderMap, pageNumbers := readLines(input)

	_, badPages := goodPagesBadPages(orderMap, pageNumbers)

	for _, bp := range badPages {
		goodPages = append(goodPages, swapElements(bp, orderMap))
	}

	for _, gp := range goodPages {
		count += middleNumber(gp)
	}
	return count
}

func main() {
	fmt.Printf("P1: %d \n", s1("input.txt"))
	fmt.Printf("P2: %d \n", s2("input.txt"))
}
