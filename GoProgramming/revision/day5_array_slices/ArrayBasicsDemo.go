package main

import "fmt"

func printArray(arr []int) {
	for index, value := range arr {
		fmt.Println("Index:", index, "Values:", value)
	}
}

func main() {
	// array declaration ways
	//var arr = [5]int{1, 2, 3, 4, 5}            // one way

	//var arr [5]int // 2nd way
	//arr[0] = 45
	//arr[3] = 44
	//arr[0] = 4000

	// copying an array into another array
	//arr1 := arr
	//arr1[3] = 55
	//
	//fmt.Println(arr1)
	//fmt.Println()
	//
	//fmt.Println(arr)
	//fmt.Println(arr[2])
	//fmt.Println(len(arr))

	//TODO: passing an array to a function
	arr := []int{000, 10, 2, 3, 4, 5, 6, 700, 8, 9}
	printArray(arr)

}
