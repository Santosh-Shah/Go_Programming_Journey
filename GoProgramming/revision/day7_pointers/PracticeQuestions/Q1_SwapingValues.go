package main

import "fmt"

func main() {
	num1 := 45
	num2 := 50

	fmt.Println("Before Swapping: ", num1, "and", num2)

	swappingValues(&num1, &num2)

	fmt.Println("After Swapping: ", num1, "and", num2)
}

func swappingValues(num1, num2 *int) {
	temp := *num1
	*num1 = *num2
	*num2 = temp
}
