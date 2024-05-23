package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a valid digit number: ")
	fmt.Scan(&num)

	output := sumOfDigit(num)
	fmt.Println("The sum of digits is:", output)
}

func sumOfDigit(num int) int {
	// Base case
	if num <= 0 {
		return 0
	}

	sum := 0
	digit := 0

	for num > 0 {
		digit = num % 10
		sum = sum + digit
		num = num / 10
	}
	return sum
}
