package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter numbers to reverse the digit numbers: ")
	fmt.Scan(&num)

	fmt.Println("Original Digits: ", num)
	output := reversedDigits(num)
	fmt.Println("Reverse Digits: ", output)
}

func reversedDigits(num int) int {
	if num < 10 {
		return num
	}

	sum := 0
	for num > 0 {
		digits := num % 10
		sum = sum*10 + digits
		num = num / 10
	}
	return sum
}
