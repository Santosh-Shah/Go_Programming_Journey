package main

import "fmt"

func sum(x, y *int) int {
	return *x + *y
}

func main() {
	var a, b int
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&b)

	result := sum(&a, &b)
	fmt.Printf("The sum is: %d\n", result)
}
