package main

import (
	"bufio"
	"fmt"
	"os"
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

func parseBinary(binary []string) int64 {
	var out int64
	var bytes []string
	//packetVersion,_ := strconv.ParseInt(strings.Join(binary[0:3],""),2,64)
	//typeId,_ := strconv.ParseInt(strings.Join(binary[3:6],""),2,64)

	start := 6
	for true {
		if binary[start] == "0" {
			bytes = append(bytes, binary[start+1:start+5]...)
			break
		}
		bytes = append(bytes, binary[start+1:start+5]...)
		start += 5
	}
	out,_ = strconv.ParseInt(strings.Join(bytes,""),2,64)
	return out
}

func solveProblem(input string) int {
	var count int

	binaries := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	//lines := readLines(input)
	// 3 first
	hexaDecimal := "D2FE28"
	var binary []string
	for _, c := range hexaDecimal {
		if _, ok := binaries[string(c)]; ok {
			for _, v := range binaries[string(c)] {
				binary = append(binary, string(v))
			}
		}
	}

	fmt.Println(parseBinary(binary))
	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}