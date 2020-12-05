package main

import "fmt"

func main(){

	characterMap := map[int]int{
		0: 127,
		7: 10,
	}

	for minNum, maxNum := range characterMap {
		fmt.Println(minNum)
		fmt.Println(maxNum)
	}
}
