package main

import "fmt"

func main() {
	sliceElem := []int{600, 500, 7, 80, 1, 3}

	output := largestElement(sliceElem)
	fmt.Println("largest element: ", output)
}

func largestElement(sliceElem []int) int {
	// Base case
	if len(sliceElem) < 1 {
		return 0
	}

	largest := sliceElem[0]
	for i := 1; i < len(sliceElem); i++ {
		if sliceElem[i] > largest {
			largest = sliceElem[i]
		}
	}
	return largest
}
