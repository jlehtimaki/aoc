package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	cache = map[string]int64{}
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

func loopThrough(binary string, memNum int64) {
	valueMap := []string{"1","0"}
	if strings.Contains(binary,"X"){
		for n, x := range binary[:]{
			if string(x) == "X" {
				for _,y := range valueMap {
					bin := binary[:n]+y+binary[n+1:]
					loopThrough(bin, memNum)
				}
			}
		}
	} else {
		cache[binary] = memNum
	}
}

func solveProblem(input string) int64{
	lines := readLines(input)
	var num int64
	var mask string
	//memMap := map[string]int64{}
	maskMap := map[int]string{}
	for n, l :=  range lines {
		fmt.Printf("Doing %d/%d\n",n,len(lines))
		if strings.Contains(l, "mask") {
			// Assign mask
			mask = strings.Split(l, "= ")[1]
			maskMap = map[int]string{}

			// Check mask assignments
			for n, m := range mask[:] {
				//bit, _ := strconv.Atoi(string(m))
				maskMap[n] = string(m)
			}
			continue
		}

		re := regexp.MustCompile("[0-9]+")

		tmp := strings.Split(l, "= ")
		memAddress := re.FindAllString(tmp[0], -1)[0]
		memNumTmp,_ := strconv.Atoi(re.FindAllString(tmp[1],-1)[0])
		memNum := int64(memNumTmp)
		memAddressNum, _ := strconv.Atoi(memAddress)
		memABit := strconv.FormatInt(int64(memAddressNum), 2)

		for len(memABit) != 36 {
			memABit = "0"+memABit
		}
		for k, v := range maskMap{
			if v == "0" {
				continue
			}
			memABit = memABit[:k] + v + memABit[k+1:]
		}
		loopThrough(memABit, memNum)

	}
	fmt.Println(cache)
	for _,v := range cache{
		num = num + v
	}
	fmt.Println(num)
	return num
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}