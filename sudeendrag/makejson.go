package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("first enter")
	firstname, _ := in.ReadString('\n')
	fmt.Println("address")

	address, _ := in.ReadString('\n')
	userMap := map[string]string{"name": strings.TrimSpace(firstname), "address": strings.TrimSpace(address)}
	//fmt.Println(userMap)
	output, err := json.Marshal(userMap)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(string(output))

	}
}
