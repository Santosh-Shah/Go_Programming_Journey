package main

import "fmt"

func main() {
	arr := [5]int{4, 7, 2, 8, 1}
	fmt.Println("Original Array:")
	printArray(arr)

	fmt.Println()

	arr = reverseArray(arr)
	fmt.Println("Reverse Array:")
	printArray(arr)
}

func printArray(arr [5]int) {
	for _, elem := range arr {
		fmt.Print(elem, " ")
	}
}

// TODO: function to reverse the array
func reverseArray(arr [5]int) [5]int {
	length := len(arr)
	i := 0
	j := length - 1
	for i < j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		i++
		j--
	}

	return arr
}
