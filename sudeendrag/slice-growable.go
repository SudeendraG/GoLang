package main

import "fmt"

func main() {
	myslice := make([]int, 1, 4)
	fmt.Printf("\nlength is : %d ,\ncapacity is : %d\n", len(myslice), cap(myslice))
	for i := 0; i < 17; i++ {
		myslice = append(myslice, i)
		fmt.Printf("\ncapacity is : %d\n", cap(myslice))

	}

	//appending slice of slice

	mycoursescompleted := []string{"docker", "awS", "java", "spring boot"} // in line declaration of splice
	newcourses := []string{"AI", "ML"}
	mycoursescompleted = append(mycoursescompleted, newcourses...)
	fmt.Println(mycoursescompleted)

}
