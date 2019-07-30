package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n-------VARADIC FUNCTIONS------\n")
	fmt.Println(vardicfun(21, 2, 22, 3, 4, 5, 6, 7, 8, 12, 123213, 23))

}

func vardicfun(league ...int) int {
	best := league[0]
	for _, i := range league {
		if i < best {
			best = i
		}
	}
	return best
}
