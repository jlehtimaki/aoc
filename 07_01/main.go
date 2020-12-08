package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//light red bags contain 1 bright white bag, 2 muted yellow bags.
//dark orange bags contain 3 bright white bags, 4 muted yellow bags.
//bright white bags contain 1 shiny gold bag.
//muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
//shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
//dark olive bags contain 3 faded blue bags, 4 dotted black bags.
//vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
//faded blue bags contain no other bags.
//dotted black bags contain no other bags.

// d-o=7 v-p=11  s-g=29
var bagNumber map[string]int

func readLines(path string) map[string]map[string]int {
	bagNumber = map[string]int{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	bagArray := map[string]map[string]int{}
	var tmpMap map[string]int
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`contain|,`)
	var bagName string
	var insideBag string
	var numBags int
	for scanner.Scan() {
		for n, s := range re.Split(scanner.Text(), -1) {
			tmpMap = map[string]int{}
			reReplace := regexp.MustCompile(`bags|bag|\.`)
			s = reReplace.ReplaceAllString(s, "")
			if n == 0 {
				bagName = strings.TrimSpace(strings.Replace(s, " ", "-", 1))
				bagArray[bagName] = map[string]int{}
				notWanted, _ := regexp.MatchString(`contain no other`, scanner.Text())
				if notWanted {
					bagNumber[bagName] = 1
				}
				continue
			}
			tmpString := strings.Split(s, " ")
			numBags,_ = strconv.Atoi(tmpString[1])
			insideBag = strings.TrimSpace(fmt.Sprintf("%s-%s", tmpString[2], tmpString[3]))

			tmpMap[insideBag] = numBags
			bagArray[bagName][insideBag] = numBags
		}
	}
	return bagArray
}

func allHaveNumbers(bag map[string]int) bool {
	for bb, _ := range bag {
		if _, ok := bagNumber[bb]; ok {
			continue
		}
		return false
	}
	return true
}

func getBags(input string) int{
	var count int
	var wantedNumber int
	bags := readLines(input)
	for len(bagNumber) != len(bags) {
		for b,_ := range bags {
			if allHaveNumbers(bags[b]){
				tmpNum := 0
				for bb,bv := range bags[b]{
					tmpNum = tmpNum + (bv * bagNumber[bb])
				}
				bagNumber[b] = tmpNum
				//fmt.Printf("Saving: %s - %d (%d - %d)\n",b, tmpNum, len(bagNumber), len(bags))
			}
		}
	}

	wantedNumber = bagNumber["shiny-gold"]
	delete(bagNumber, "shiny-gold")

	for _, v := range bagNumber {
		if v >= wantedNumber{
			count++
		}
	}
	return count
}

func main()  {
	lines := getBags("input.txt")
	fmt.Println(lines)
}
