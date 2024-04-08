package main

import "fmt"

func mul(a1, a2 int) int {

	res := a1 * a2
	fmt.Println("Result: ", res)
	return 0
}

func show() {
	fmt.Println("Hello!, GeeksforGeeks")
}

// Main function
func main() {

	// Calling mul() function
	// Here mul function behaves
	// like a normal function
	//TODO:
	//mul(23, 45)

	// Calling mul()function
	// Using defer keyword
	// Here the mul() function
	// is defer function
	//TODO:
	//defer mul(23, 10)

	// Calling show() function
	//TODO:
	//show()

	//TODO: creating multiple defer methods
	fmt.Println("Start")
	defer fmt.Println("End")
	defer add(24, 11)
	defer add(20, 5)
}

func add(v1 int, v2 int) {
	fmt.Println("Result: ", v1+v2)
}
