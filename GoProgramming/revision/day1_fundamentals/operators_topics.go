package main

import "fmt"

func main() {
	var val1 int = 30
	var val2 int = 10

	fmt.Println("Addition:", val1+val2)
	fmt.Println("Subtraction:", val1-val2)
	fmt.Println("Multiplication:", val1*val2)
	fmt.Println("Division:", val1/val2)
	fmt.Println("Modulo:", val1%val2)

	fmt.Println("Equal:", val1 == val2)
	fmt.Println("Not Equal:", val1 != val2)
	fmt.Println("Greater than:", val1 > val2)
	fmt.Println("val1:", val1)
	fmt.Println("val2:", val2)
}
