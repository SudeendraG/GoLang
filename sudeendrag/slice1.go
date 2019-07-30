package main

import "fmt"

func main() {
	// declaration of slice
	mycourse := make([]string, 5, 10)
	mycoursescompleted := []string{"docker", "awS", "java", "spring boot"} // in line declaration of slice
	fmt.Printf("length is : %d ,\ncapacity is : %d\n", len(mycourse), cap(mycourse))
	fmt.Println(mycoursescompleted[2])
	mycoursescompleted[2] = "java 1.8"
	fmt.Println(mycoursescompleted)

	// sliceOf Slice

	sliceofSclice := mycoursescompleted[2:4] //exclude 4 , include 2
	fmt.Println(sliceofSclice)

	sliceofSclice = mycoursescompleted[:4] // considerd as 0:4
	fmt.Println(sliceofSclice)

	sliceofSclice = mycoursescompleted[4:] //4:length-1
	fmt.Println(sliceofSclice)

}
