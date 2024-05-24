package main

import "fmt"

func main() {
	array := []int{0, 205, 2, 100}
	output := findMaxElement(array)
	fmt.Println("Maximum element in array:", output)
}

func findMaxElement(array []int) int {
	if len(array) <= 0 {
		return -1
	}

	max := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}

	return max
}
