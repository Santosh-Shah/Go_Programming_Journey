package main

import "fmt"

func main() {
	arr := []int{-1, -6, 7, 9, 1, 10, 45}

	// it is declared in same package in another files
	fmt.Println("Original Array: ")
	printSliceElement(arr)

	fmt.Println()

	fmt.Println("After Bubble Sor: ")
	bubbleSort(arr)
	printSliceElement(arr)
}

// function for bubble sort
func bubbleSort(arr []int) {
	// Base case
	if len(arr) <= 0 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func printSliceElement(arr []int) {
	for _, elem := range arr {
		fmt.Print(elem, " ")
	}
}
