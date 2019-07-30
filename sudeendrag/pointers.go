package main

import (
	"fmt"
)

func main() {
	//function level
	module := 32
	ptr := &module
	fmt.Println("\nMemory address  of **module", &module)
	fmt.Println("Memory address  of **module", &module, "and value of module is ", *ptr)

	//pass by value
	fmt.Println("\n------------------PASS BY VALUE------------\n")
	course := "golang"
	fmt.Println("value of Course is", course)
	changeCourse(course)
	fmt.Println("Final course value", course)

	//pass by refernce
	fmt.Println("\n------------------PASS BY Reference------------\n")

	course2 := "golang"
	fmt.Println("value of Course2 is", course2)
	changeCourse2(&course2)
	fmt.Println("Final course2 value", course2)
}

func changeCourse(course string) string {

	course = "java"
	fmt.Println("the updated value of course is ", course)
	return course
}

func changeCourse2(course2 *string) string {

	*course2 = "java"
	fmt.Println("the updated value of course2 is ", *course2)
	return *course2
}
