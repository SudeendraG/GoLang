package main

import "fmt"

var input float64

func main() {

	//fmt.Println("enter a floating point number")
	fmt.Scan(&input)
	trunc := int64(input)
	//fmt.Println("truncated version of the floating point number is", trunc)
	fmt.Println(trunc)
}
