package main

import "fmt"

func main() {
	//TODO: Arithmetic Operators
	x := 10
	y := 5

	sum := x + y
	difference := x - y
	product := x * y
	quotient := x / y
	remainder := x % y
	fmt.Println("Arithmetic Operations")
	fmt.Println("Sum: ", sum)
	fmt.Println("Difference: ", difference)
	fmt.Println("Product: ", product)
	fmt.Println("Quotient: ", quotient)
	fmt.Println("Remainder: ", remainder)

	//TODO: Relational Operators
	isEqual := x == y
	isNotEqual := x != y
	isLessThan := x < y
	isGreaterThan := x > y
	isLessThanOrEqual := x <= y
	isGreaterThanOrEqual := x >= y
	fmt.Println("\nRelational Operations:")
	fmt.Println("Equal:", isEqual)
	fmt.Println("Not equal:", isNotEqual)
	fmt.Println("Less than:", isLessThan)
	fmt.Println("Greater than:", isGreaterThan)
	fmt.Println("Less than or equal:", isLessThanOrEqual)
	fmt.Println("Greater than or equal:", isGreaterThanOrEqual)

	//TODO: logical Operators
	isPositive := x > 0 && y > 0
	isEvenOrFive := x%2 == 0 || y == 5
	isNotZero := !(x == 0)
	fmt.Println("\nLogical Operations:")
	fmt.Println("Positive:", isPositive)
	fmt.Println("Even or five:", isEvenOrFive)
	fmt.Println("Not zero:", isNotZero)

	//TODO: Bitwise Operators
	a := 10
	b := 6
	andResult := a & b
	orResult := a | b
	xorResult := a ^ b
	leftShift := a << 2
	rightShift := a >> 1
	fmt.Println("\nBitwise Operation:")
	fmt.Println("And:", andResult)          // Output: 2
	fmt.Println("Or:", orResult)            // Output: 14
	fmt.Println("Xor:", xorResult)          // Output: 12
	fmt.Println("Left shift:", leftShift)   // Output: 40
	fmt.Println("Right shift:", rightShift) // Output: 5

	//TODO: Assignment operator
	var num = 20
	num += 5
	num -= 3
	fmt.Println("\nAssignment: ")
	fmt.Println("Number: ", num)

	// Misc operators
	pointer := &num   // address-of operator, gets the memory address of num
	value := *pointer // dereference operator, accesses the value stored at the memory address pointed to by pointer
	fmt.Println("\nMisc:")
	fmt.Println("Pointer:", pointer)          // Output: memory address of num (varies depending on execution)
	fmt.Println("Value from pointer:", value) // Output: 22
}
