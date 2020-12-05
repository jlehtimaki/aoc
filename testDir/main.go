package main

import (
	"fmt"
	"regexp"
)

func main(){

	string := "hzl"
	eclTrue, _ := regexp.MatchString(`^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$`, string)
	if eclTrue == false {
		fmt.Println("false")
	}
}
