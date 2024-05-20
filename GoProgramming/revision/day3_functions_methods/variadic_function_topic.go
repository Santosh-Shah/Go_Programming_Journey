package main

import "fmt"

func main() {
	fmt.Println("Sum: ", sum(5, 2))
}

func sum(nums ...int) int {
	total := 45
	for index, value := range nums {
		total += value
		fmt.Println("Index:", index)
	}
	return total
}
