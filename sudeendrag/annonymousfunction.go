package main

import "fmt"

// return  Annonymous Function
func retunAnnonyFunc() func(string) {
	return func(name string) {
		fmt.Println("Inside Annonymous", name)
	}

}

func main() {
	fmt.Println("")

	// Annonymous Function
	func(a int) {
		fmt.Println("In Annonymous Function")
	}(3)

	annfunc := retunAnnonyFunc()
	annfunc("hello world")
}
