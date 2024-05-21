package main

import (
	"fmt"
)

func modifySlice(slice []int) {
	slice[4] = 400000
}

func areSlicesEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}

func main() {
	//TODO: Creating a slice from an array
	//arr := [5]int{1, 2, 3, 4, 5}
	//slice := arr[1:4] //Slice from index 1 to 3
	//fmt.Println(slice)

	//TODO: Slice Composite Literal
	// means slice can directly declare without creating array explicitly
	//slice := []int{10, 20, 30, 4, 5}
	//fmt.Println(slice)

	//TODO: Copying one slice into another slice
	//slice1 := []int{111, 2222, 3333, 4444}
	//slice2 := make([]int, len(slice1))
	////slice2 := make([]int, 10)
	//copy(slice2, slice1)
	//
	//fmt.Println("Slice 1:", slice1)
	//fmt.Println("Slice 2:", slice2)

	//TODO: passing a slice to function
	// slice is a pass by reference means,
	// if we change the value of slice inside function then
	// actual value of slice is also changes

	//slice := []int{12, 120, 45, 333, 68}
	//
	//fmt.Println("Original slice:", slice)
	//
	//modifySlice(slice)
	//fmt.Println("New slice:", slice)

	//TODO: Comparing slices with function and without function
	slice1 := []int{12, 13, 14, 15}
	slice2 := []int{22, 33, 44}
	slice3 := []int{12, 13, 14, 15}

	// Direct comparing
	//fmt.Println(slices.Equal(slice1, slice2))
	//fmt.Println(slices.Equal(slice1, slice3))
	//
	//fmt.Println(len(slice1) == len(slice2))
	//fmt.Println(len(slice1) == len(slice3))

	// with function
	fmt.Println("Slice1 equals Slice2:", areSlicesEqual(slice1, slice2))
	fmt.Println("Slice1 equals Slice3:", areSlicesEqual(slice1, slice3))
}
