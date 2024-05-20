package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}

	result := add(45, 12)
	fmt.Println("Result:", result)
}
