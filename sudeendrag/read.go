package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var filename string
	fmt.Println("enter file name")
	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		panic(err)
	} else if scanner.Scan() {
		filename = scanner.Text()
	}

	// open file
	file, err := os.Open(filename)
	if err == nil {
		// build struct
		constructStruct(file)
	} else {
		fmt.Println(err)
	}
	defer file.Close()

}

type User struct {
	firstname string
	lastname  string
}

func constructStruct(file io.Reader) {
	arr := []*User{}
	scanner := bufio.NewScanner(file)
	fmt.Println("")
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		data := new(User)
		data.firstname = words[0]
		data.lastname = words[1]
		arr = append(arr, data)
	}
	//loop slice of struts
	for _, data := range arr {
		fmt.Println(data.firstname, data.lastname)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
