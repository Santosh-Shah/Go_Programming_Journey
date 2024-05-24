package main

import "fmt"

func main() {
	array := []int{4, 5, 6, 2, 3}

	output := sumOfArrayElement(array)
	fmt.Println("Sum of array elements:", output)
}

func sumOfArrayElement(array []int) int {
	// Base case
	if len(array) < 1 {
		return -1
	}

	sum := 0
	i := 0
	for i != len(array) {
		sum += array[i]
		i++
	}
	return sum
}
