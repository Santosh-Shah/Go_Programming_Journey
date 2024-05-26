package main

import "fmt"

func main() {
	//arr := []int{4, 5, 3, 2, 10}
	//output := sumOfElements(arr)
	//fmt.Println("Sum of all elements: ", output)

	// taking input from user
	var length int
	fmt.Print("Enter the length of the array: ")
	fmt.Scan(&length)

	// Create an array
	array := make([]int, length)

	// taking input from user
	fmt.Println("Enter the elements of the array: ")
	for i := 0; i < length; i++ {
		fmt.Printf("Element %d: ", i)
		fmt.Scan(&array[i])
	}

	// calculating the sum
	sum := sumOfElements(array)
	fmt.Println("The sum of all elements is: ", sum)
}

func sumOfElements(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}
