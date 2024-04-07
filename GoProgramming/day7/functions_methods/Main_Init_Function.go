package main

import "fmt"

func main() {
	//s := []int{235, 34, 1, 3, 70}
	//sort.Ints(s)
	//fmt.Println("Sort slice: ", s)
	//
	//// find the index
	//fmt.Println("Index value: ", strings.Index("Vedas College", "V"))
	//
	////finding the time
	//fmt.Println("Time: ", time.Now())

	//Init function demo

	fmt.Println("Welcome to main() function")

}

func init() {
	fmt.Println("Welcome to init()1 function")
}

func init() {
	fmt.Println("Welcome to init()2 function")
}
