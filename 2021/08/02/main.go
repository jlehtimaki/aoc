package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func sortStringArray(sl []string) []string {
	var rt []string
	for _, s := range sl {
		rt = append(rt, sortString(s))
	}
	return rt
}

func sortString(s string) string {
	cl := strings.Split(s, "")
	sort.Strings(cl)
	cs := strings.Join(cl,"")
	return cs
}

func stringContains(x string, y string) bool {
	if len(y) < 2 {
		return false
	}
	for _, c := range y {
		if !strings.Contains(x, string(c)) {
			return false
		}
	}
	return true
}

func determineNumbers(sa []string) map[string]string {
	numbers := map[string]string{}
	for true {
		for _, s := range sa {
			s = sortString(s)
			if len(s) == 2 {
				numbers["1"] = s
				continue
			}
			if len(s) == 4 {
				numbers["4"] = s
				continue
			}
			if len(s) == 3 {
				numbers["7"] = s
				continue
			}
			if len(s) == 7 {
				numbers["8"] = s
				continue
			}
			if len(s) == 5 {
				// 2, 3, 5,
				if stringContains(s, numbers["1"]){
					numbers["3"] = s
					continue
				}
				if _, ok := numbers["1"]; ok {
					f := strings.Replace(numbers["4"], string(numbers["1"][0]), "", -1)
					f = strings.Replace(f, string(numbers["1"][1]),"", -1)
					if stringContains(s, f) {
						numbers["5"] = s
						continue
					}
				}
				numbers["2"] = s
			}
			if len(s) == 6 {
				// 0, 6, 9
				if stringContains(s, numbers["4"]) && stringContains(s, numbers["3"]){
					numbers["9"] = s
					continue
				}
				if stringContains(s, numbers["7"]){
					if _, ok := numbers["9"]; ok {
						numbers["0"] = s
						continue
					}
				}
				numbers["6"] = s
			}
		}
		fmt.Println(numbers)
		if len(numbers) == 10 {
			break
		}
	}

	return numbers
}

func solveProblem(input string) int {
	var count int
	lines := readLines(input)

	for _, line := range lines {
		var countString string
		lineArray := strings.Split(line, " | ")
		afterPart := lineArray[1]
		afSplit := strings.Split(afterPart, " ")
		afSplit = sortStringArray(afSplit)
		numbers := determineNumbers(strings.Split(lineArray[0], " "))
		fmt.Println(numbers)
		fmt.Println(afSplit)
		for _, a := range afSplit {
			for k,v := range numbers {
				if a == v {
					countString = countString + k
					break
				}
			}
		}
		//fmt.Println(countString)
		countInt,_ := strconv.Atoi(countString)
		count = count + countInt
	}

	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}