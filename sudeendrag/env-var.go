package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USER")
	fmt.Println("\n", os.Environ())
	fmt.Println("\nLogged in user is : ", name)
}
