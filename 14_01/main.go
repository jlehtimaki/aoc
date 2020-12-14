package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func solveProblem(input string) int64{
	lines := readLines(input)
	var num int64
	var mask string
	memMap := map[string]int64{}
	maskMap := map[int]string{}
	for _, l :=  range lines {
		if strings.Contains(l, "mask"){
			// Assign mask
			mask = strings.Split(l, "= ")[1]
			maskMap = map[int]string{}

			// Check mask assignments
			for n, m := range mask[:] {
				if string(m) == "X" {
					continue
				}
				//bit, _ := strconv.Atoi(string(m))
				maskMap[n] = string(m)
			}
			continue
		}

		re := regexp.MustCompile("[0-9]+")

		tmp := strings.Split(l, "= ")
		memAddress := re.FindAllString(tmp[0], -1)[0]
		memInt,_ := strconv.Atoi(re.FindAllString(tmp[1], -1)[0])
		memBit := strconv.FormatInt(int64(memInt), 2)
		//fmt.Printf("Address: %s - MemoryInt: %d - MemoryBit: %s ",memAddress, memInt, memBit)

		// Convert memInt to 36bit string
		for len(memBit) != 36 {
			memBit = "0" + memBit
		}

		memArray := []string{}
		for _, ms := range memBit{
			memArray = append(memArray, string(ms))
		}
		//fmt.Printf(" - MemArray: %s - ", memArray)
		for k,v := range maskMap {
			memArray[k] = v
		}
		//fmt.Printf("MaskMap: %v - ", maskMap)
		memBit = strings.Join(memArray,"")
		//fmt.Printf("Membit: %s \n", memBit)

		mInt,_ := strconv.ParseInt(memBit, 2, 64)
		memMap[memAddress] = mInt
		fmt.Println(memMap)
	}
	var s int
	for _, v := range memMap{
		num = num + v
		s++
	}
	return num
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}