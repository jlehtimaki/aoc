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

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func arrayToMap(array []string) map[string]int {
	m := map[string]int{}
	for _, a := range array {
		m[a] = 0
	}
	return m
}

func solveProblem(input string) int {
	lines := readLines(input)
	e := arrayToMap(lines)
	g := arrayToMap(lines)
	var eeString string
	var ggString string

	for i:=0; i < len(lines[0]); i++ {
		if len(g) > 1 {
			var gString string
			var kv string
			for k, _ := range g {
				gString = gString + string(k[i])
			}

			if strings.Count(gString, "1") > (len(g) / 2) || strings.Count(gString, "1") == strings.Count(gString, "0"){
				kv = "1"
			} else {
				kv = "0"
			}

			for k,_ := range g {
				if string(k[i]) != kv {
					delete(g, k)
				}
			}
		}

		if len(e) > 1 {
			var eString string
			var kv string
			for k, _ := range e {
				eString = eString + string(k[i])
			}

			if strings.Count(eString, "1") > (len(e) / 2) || strings.Count(eString, "1") == strings.Count(eString, "0") {
				kv = "0"
			} else {
				kv = "1"
			}

			for k,_ := range e {
				if string(k[i]) != kv {
					delete(e, k)
				}
			}
		}
	}
	//fmt.Println(bitArray)
	//fmt.Println(e)
	//fmt.Println(g)

	for k, _ := range g {
		ggString = k
	}
	for k, _ := range e {
		eeString = k
	}
	eInt,_ := strconv.ParseInt(eeString,2, 64)
	gInt,_ := strconv.ParseInt(ggString, 2, 64)


	return int(eInt * gInt)
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}