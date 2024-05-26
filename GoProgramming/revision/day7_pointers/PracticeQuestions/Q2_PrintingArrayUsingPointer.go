package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array elements: ")
	printArrayUsingPointer(&array)
}

func printArrayUsingPointer(array *[5]int) {
	for _, value := range *array {
		fmt.Print(value, " ")
	}
}
