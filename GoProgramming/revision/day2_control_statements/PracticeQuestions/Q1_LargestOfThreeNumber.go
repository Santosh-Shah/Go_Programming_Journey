package main

import (
	"fmt"
)

func largestOfThreeNumber(n1, n2, n3 int) int {
	if n1 < 0 || n2 < 0 || n3 < 0 {
		return -1
	}

	if n1 >= n2 && n1 >= n3 {
		return n1
	} else if n2 >= n3 {
		return n2
	} else {
		return n3
	}
}

func main() {

	var n1, n2, n3 int
	fmt.Println("Enter an three integer number:")
	fmt.Scanln(&n1, &n2, &n3)
	output := largestOfThreeNumber(n1, n2, n3)
	fmt.Println("Largest of number is:", output)
}
