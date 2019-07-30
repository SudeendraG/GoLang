package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arr := make([]int, 0, 3)
	input := ""
	for i := 0; ; i++ {
		fmt.Println("enter the integer")
		fmt.Scan(&input)
		if strings.EqualFold(strings.ToLower(input), "x") {
			break
		}
		value, _ := strconv.ParseInt(input, 10, 64)
		arr = append(arr, int(value))
		//sorting
		sort.Ints(arr)
		fmt.Println(arr)
	}
}
