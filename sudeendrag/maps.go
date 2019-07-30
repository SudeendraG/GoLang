package main

import "fmt"

func main() {

	leagugetitles := make(map[string]int)
	leagugetitles["Manchister"] = 6
	leagugetitles["sunderland"] = 3

	// Inline Map Initialization
	recentHead2Head := map[string]int{
		"suderland": 5,
		"newcastle": 3,
	}

	fmt.Printf("\n leagueTitles : %v\n recenthead2head : %v\n", leagugetitles, recentHead2Head)

	//iterating map

	testmap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	for k, v := range testmap {
		fmt.Println("key : ", k, " value : ", v)
	}

	//manupulating map elements
	fmt.Println("")
	fmt.Println(testmap)
	//update
	testmap["A"] = 1000
	fmt.Println(testmap)

	// insert new
	testmap["Z"] = 1212
	fmt.Println(testmap)

	// delete element in map
	delete(testmap, "Z")
	fmt.Println(testmap)

}
