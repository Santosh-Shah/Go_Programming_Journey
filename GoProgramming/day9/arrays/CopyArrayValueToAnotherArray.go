package main

import "fmt"

func main() {
	//arr1 := [4]string{"maneesh", "priya", "ravi", "rohit"}
	//
	//// coping in another variable
	//arr2 := arr1
	//fmt.Println("Array 1:", arr1)
	//fmt.Println("Array 2:", arr2)
	//
	//arr1[0] = "Santosh"
	//fmt.Println("Array 1:", arr1)
	//fmt.Println("Array 2:", arr2)

	//TODO: same coping implementation but with address
	//arr1 := [4]string{"maneesh", "priya", "ravi", "rohit"}
	//
	//arr2 := &arr1
	//
	//fmt.Println("Array 1:", arr1)
	//fmt.Println("Array 2:", *arr2)
	//
	//arr1[0] = "Santosh"
	//fmt.Println("Array 1:", arr1)
	//fmt.Println("Array 2:", *arr2)

	//TODO: another way to copy array
	//originalArray := []int{1, 2, 3, 4, 5}
	//copyArray := make([]int, 0, len(originalArray))
	//
	//copyArray = append(copyArray, originalArray...)
	//
	//fmt.Println("Original Array: ", originalArray)
	//fmt.Println("Copy Array: ", copyArray)

	//TODO: passing array through function
	var arr = [6]int{23, 45, 33, 22, 33, 66}
	//var res int

	// passing an array as an argument
	value := myfun(arr, len(arr))
	fmt.Printf("Final result: %d ", value)

}

func myfun(arr [6]int, size int) int {
	var value int

	for i := 0; i < size; i++ {
		value = value + arr[i]
	}
	return value
}
