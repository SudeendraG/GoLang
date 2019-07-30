package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	input, _ := in.ReadString('\n')
	//fmt.Println(input)
	// convert everyting to lowercase
	//input = strings.ToLower(input)
	input = strings.ToLower(strings.TrimSpace(input))

	if strings.EqualFold(input[0:1], "i") && strings.EqualFold(input[len(input)-1:], "n") && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("“Not Found!")
	}

}
