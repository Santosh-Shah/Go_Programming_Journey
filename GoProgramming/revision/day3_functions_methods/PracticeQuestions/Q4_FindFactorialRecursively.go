package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number to calculate factorial Recursively: ")
	fmt.Scan(&num)

	output := findFactorial(num)
	fmt.Printf("The factorial of %d Recursively is: %d", num, output)
}

func findFactorial(num int) int {
	if num == 0 {
		return 1
	}

	output := findFactorial(num - 1)
	return output * num
}
