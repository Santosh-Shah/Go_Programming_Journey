package main

import "fmt"

func main() {
	var value int = 100

	var ptr1 *int = &value

	var ptr2 **int = &ptr1

	fmt.Println("Value of value: ", value)
	fmt.Println("Address of value: ", &value)

	fmt.Println("Value of ptr1: ", *ptr1)
	*ptr1 = 14
	fmt.Println("after changing Value of ptr1: ", *ptr1)
	fmt.Println("Address of pt1: ", ptr1)

	fmt.Println("Address of ptr2: ", *ptr2)
	fmt.Println("Value of ptr2: ", **ptr2)

	**ptr2 = 1444
	fmt.Println("After changing Value of ptr2: ", **ptr2)
}
