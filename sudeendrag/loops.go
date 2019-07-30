package main

import (
	"fmt"
	"time"
)

func main() {

	// for loop
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Bhoom!!!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)

	}
	fmt.Println("\n ------- * For Range Loop * ------- \n")
	// for range loop
	courseCatalog := []string{"Docker deep dive", "Neural Networks", "Deep learning", "Machine learning", "Data Science", "AI"}
	for index, data := range courseCatalog { // for _,i := range cousecatalog {}
		fmt.Println(index)
		fmt.Println(data)
	}
}
