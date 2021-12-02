package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readLines(path string) []map[string]string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var stringArray []map[string]string
	tmpString := map[string]string{}
 	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			splitString := strings.Split(scanner.Text()," ")
			for _, sp := range splitString {
				s := strings.Split(sp, ":")
				tmpString[s[0]] = s[1]
			}
			continue
		}
		stringArray = append(stringArray, tmpString)
		tmpString = map[string]string{}
	}
	stringArray = append(stringArray, tmpString)
	return stringArray
}

type Passport struct {
	byr string //(Birth Year)
	iyr string //(Issue Year)
	eyr string //(Expiration Year)
	hgt string //(Height)
	hcl string //(Hair Color)
	ecl string //(Eye Color)
	pid string //(Passport ID)
	cid string //(Country ID)
}

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	//"cid",
}

func checkRequiredFields(passport map[string]string) bool {
	for _, x := range requiredFields {
		if passport[x] == "" {
			return false
		}
	}
	return true
}

func main(){

	//passport := Passport{}

	file := readLines("input.txt")
	var count int
	for _, f := range file {
		if checkRequiredFields(f) {
			count++
		}
	}
	fmt.Println(count)
}
