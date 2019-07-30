package main

import (
	"fmt"
	"reflect"
)

// package level
var x int = 98
var (
	name string
	id   int
)

func main() {
	//function level
	var i int
	i = 90
	fmt.Println("hi this is go", x)
	fmt.Println(i)
	// onother way of defining
	z := 290.989
	fmt.Println(z)
	var complexnumber complex64
	fmt.Println(complexnumber)
	fmt.Println("Type of z", reflect.TypeOf(z))
	fmt.Println("Type of complexnumber", reflect.TypeOf(complexnumber))
	fmt.Println("Type of i", reflect.TypeOf(i))

}
