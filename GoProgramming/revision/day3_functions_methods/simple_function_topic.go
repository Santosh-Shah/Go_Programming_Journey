package main

import "fmt"

func main() {
	result := add(40, 5)
	fmt.Println("Result:", result)
}

func add(x, y int) int {
	return x + y
}
