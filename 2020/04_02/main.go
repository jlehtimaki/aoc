package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

type Passport struct {
	byr int //(Birth Year)
	iyr int //(Issue Year)
	eyr int //(Expiration Year)
	hgt string //(Height)
	hcl string //(Hair Color)
	ecl string //(Eye Color)
	pid string //(Passport ID)
	cid string //(Country ID)
}
//map[byr:1992 cid:277 ecl:brn eyr:2020 hcl:dab227 hgt:182cm iyr:2012 pid:021572410]

//byr (Birth Year) - four digits; at least 1920 and at most 2002.
//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
//hgt (Height) - a number followed by either cm or in:
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
//pid (Passport ID) - a nine-digit number, including leading zeroes.
//cid (Country ID) - ignored, missing or not.

func checkRequiredFields(passport map[string]string) bool {
	for _, x := range requiredFields {
		if passport[x] == "" {
			return false
		}
	}

	byr,_ := strconv.Atoi(passport["byr"])
	iyr,_ := strconv.Atoi(passport["iyr"])
	eyr,_ := strconv.Atoi(passport["eyr"])
	//pid,_ := strconv.Atoi(passport["pid"])

	p := Passport{
		byr: byr,
		iyr: iyr,
		eyr: eyr,
		hgt: passport["hgt"],
		hcl: passport["hcl"],
		ecl: passport["ecl"],
		pid: passport["pid"],
	}

	if p.byr < 1920 || p.byr > 2002 {
		return false
	}

	if p.iyr < 2010 || p.iyr > 2020 {
		return false
	}

	if p.eyr < 2020 || p.eyr > 2030 {
		return false
	}

	hgtTrue, _ := regexp.MatchString(`cm|in`, p.hgt)
	if hgtTrue == true {
		re := regexp.MustCompile("[0-9]+")
		height, _ := strconv.Atoi(re.FindString(p.hgt))

		if strings.Contains(p.hgt, "in") {
			if height < 59 || height > 76 {
				return false
			}
		} else {
			if height < 150 || height > 193 {
				return false
			}
		}
	} else {
		return false
	}

	hclTrue, _ := regexp.MatchString(`^\#[a-f0-9]{6,}$`, p.hcl)
	if hclTrue == false {
		return false
	}

	eclTrue, _ := regexp.MatchString(`^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$`, p.ecl)
	if eclTrue == false {
		return false
	}

	pidTrue, _ :=  regexp.MatchString(`^\d{9}$`, p.pid)
	if pidTrue == false {
		return false
	}

	return true
}

func main(){

	file := readLines("input.txt")
	var count int
	for _, f := range file {
		if checkRequiredFields(f) {
			fmt.Println(f)
			count++
		}
	}
	fmt.Println(count)
}
