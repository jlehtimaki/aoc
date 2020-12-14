package main

import "fmt"

func main(){
	foobar := "X1011100000X111X01001000001110X00000"
	z := "11111101101001110001001"
	for len(z) != len(foobar){
		z = "0"+z
	}
	fmt.Println(foobar)
	fmt.Println(z)
}
