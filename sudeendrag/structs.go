package main

import "fmt"

func main() {

	type courseMeta struct {
		auther string
		level  string
		rating float32
	}

	//var docker courseMeta
	//docker := new(courseMeta)

	//litral way

	docker := courseMeta{
		auther: "Paul",
		level:  "Intermediate",
		rating: 3,
	}

	fmt.Println(docker)
	fmt.Println("Auther : ", docker.auther)
	fmt.Println("rating : ", docker.rating)
}
