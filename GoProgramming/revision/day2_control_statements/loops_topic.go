package main

import "fmt"

func main() {
	//TODO: basic loop
	//for i := 0; i < 5; i++ {
	//	fmt.Println("Vedas College:", i)
	//}

	//TODO: for loop as while loop
	//i := 0
	//for i < 5 {
	//	fmt.Println("Vedas College:", i)
	//	i++
	//}

	//TODO: Infinite for loop
	i := 0
	for {
		fmt.Println("Vedas College", i)
		i++
		if i == 5 {
			break
		}
	}
}
