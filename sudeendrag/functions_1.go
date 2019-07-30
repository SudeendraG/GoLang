package main

import (
	"fmt"
	"strings"
)

func main() {

	auther, module := converter("msdhoni", "cricket")
	fmt.Println(auther, module)

}
func converter(auther, module string) (a, m string) {

	module = strings.ToUpper(module)
	auther = strings.ToUpper(auther)
	return auther, module
}
