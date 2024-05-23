package main

import "fmt"

func findFactorial(num int) int {
	// Base case
	if num <= 1 {
		return num
	}

	factorial := 1
	for i := num; i >= 1; i-- {
		factorial *= i
	}
	return factorial
}

func main() {
	var num int
	fmt.Println("Enter a number: ")
	fmt.Scan(&num)
	fmt.Println("Factorial of", num, "is:", findFactorial(num))
}
