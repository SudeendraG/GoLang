package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("\n-----------Conditionals-----------------  \n")

	//switch case
	switch "docker" {
	case "linux":
		fmt.Println("\n Linux course \n")
	case "docker":
		fmt.Println("\n docker course \n")
	case "windows":
		fmt.Println("\n Windows course \n")
	default:
		fmt.Println("\n no default fall back course \n")

	}

	// switch case from expression

	switch temp := random(); temp {
	case 0, 2, 4, 6, 8:
		fmt.Println("\n even \n")
	case 1, 3, 5, 7, 9:
		fmt.Println("\n odd \n")
	}

}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)

}
