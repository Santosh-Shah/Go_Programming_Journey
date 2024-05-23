package main

import "fmt"

func evenOrOdd(num int) {
	// Base case
	if num <= 1 {
		fmt.Println(num, ": number must be equal or greater than 2")
		return
	}

	if num%2 == 0 {
		fmt.Println(num, ":is even number")
	} else {
		fmt.Println(num, ":is odd number")
	}
}

func main() {
	var num int
	fmt.Println("Enter a number equal to or greater than 2: ")
	fmt.Scan(&num)
	evenOrOdd(num)
}
