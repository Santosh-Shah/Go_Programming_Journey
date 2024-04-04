package main

import "fmt"

func main() {
	x := 500
	fmt.Println("Value of X:", x)

	//changing the value of x by pointer function
	//ptr := &x
	//ptrF(ptr)

	// other ways to do this same is
	ptrF(&x)
	fmt.Println("New value of x: ", x)
}

func ptrF(newValue *int) {
	*newValue = 1000
}
