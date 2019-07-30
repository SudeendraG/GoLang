package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("\n-----------Conditionals-----------------  \n")
	firstrank := "39"
	secondrank := "614"
	// if firstrank , secondrank := 39,512 ; firstrank < secondrank {    // can also write like this
	if firstrank < secondrank {
		fmt.Println("\n first course is doing better \n")
	} else if firstrank > secondrank {
		fmt.Println("\n second course is doing better \n")
	} else {
		fmt.Println("\n niether first or second is doing better \n")
	}

	//role of if in error handling
	_, err := os.Open("/Users/sudeendra.g/Documents/index1.html")
	if err != nil {
		fmt.Println("error occurd i.e, ", err)
	}
	_, err2 := os.Open("/Users/sudeendra.g/Documents/index.html")
	fmt.Println("nill error  ", err2)

}
