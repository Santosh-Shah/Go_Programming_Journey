package main

import (
	"fmt"
)

func main() {

	//TODO: Soring slice
	slice := []int{4, 5, 7, 1, 23, 0}
	//fmt.Println(slice)
	//sort.Ints(slice)
	//fmt.Println(slice)

	//TODO: Trimmming a slice
	//trimmedSlice := slice[1:4]
	//fmt.Println(trimmedSlice)

	//TODO: Splitting a slice
	//slice1 := slice[:2]
	//slice2 := slice[2:]
	//fmt.Println("First Part:", slice1)
	//fmt.Println("Second Part:", slice2)
	fmt.Println("Original slice:", slice)

	//start := 0
	//end := 1
	//fmt.Println(slice[start:end])

	fmt.Println(slice[:1])
}
