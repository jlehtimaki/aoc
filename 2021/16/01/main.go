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

func parsePacket(binary []string, start int) int64 {
	packetVersion,_ := strconv.ParseInt(strings.Join(binary[start:start+3],""),2,64)
	var packetLenght int
	var out int64
	packetType,_ := strconv.ParseInt(strings.Join(binary[start+3:start+6],""),2,64)
	packetTypeId := strings.Join(binary[6:7],"")

	fmt.Println(packetVersion)
	fmt.Println(packetType)
	fmt.Println(packetTypeId)
	if packetType == 4 {
		return solveBinary(binary)
	} else {
		if packetTypeId == "0" {
			var version int64
			packetLenght = 15
			subPackets, _ := strconv.ParseInt(strings.Join(binary[7:7+packetLenght],""),2,64)
			start := 7+packetLenght
			var end int
			for out < subPackets {
				firstBin := strings.Join(binary[start:start+1],"")
				if firstBin == "1" {
					end = start + 11

				} else if firstBin == "0" {
					end = start + 15
				}
				out = out + solveBinary(binary[start:end])
				start = end
			}
		} else if packetTypeId == "1" {
			packetLenght = 11
			subPackets, _ := strconv.ParseInt(strings.Join(binary[7:7+packetLenght],""),2,64)
			start := 7+packetLenght
			for i:=0; i<int(subPackets); i++ {
				end := start + packetLenght
				out = out + solveBinary(binary[start:start+packetLenght])
				start = end
			}
		}
	}

	return out
}


func solveBinary(binary []string) int64 {
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
	hexaDecimal := "A0016C880162017C3686B18A3D4780"
	var binary []string
	for _, c := range hexaDecimal {
		if _, ok := binaries[string(c)]; ok {
			for _, v := range binaries[string(c)] {
				binary = append(binary, string(v))
			}
		}
	}

	binary = append(binary, "0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0")
	fmt.Println(binary)
	fmt.Println(parsePacket(binary,0))
	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}