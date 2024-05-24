package main

import "fmt"

func main() {
	var length, width int
	fmt.Println("Enter length and width of rectangle: ")
	fmt.Scan(&length, &width)

	output := areaOfRectangle(length, width)
	fmt.Println("Area of rectangle is:", output)
}

func areaOfRectangle(length, width int) int {
	// Base case
	if length == 0 || width == 0 {
		return 0
	}

	return length * width
}
