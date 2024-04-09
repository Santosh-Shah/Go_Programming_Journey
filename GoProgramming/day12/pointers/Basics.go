package main

import "fmt"

func main() {
	//x := 0xFF
	//y := 0x9C
	//// Displaying the values
	//fmt.Printf("Type of variable x is %T\n", x)
	//fmt.Printf("Value of x in hexadecimal is %X\n", x)
	//fmt.Printf("Value of x in decimal is %v\n", x)
	//
	//fmt.Printf("Type of variable y is %T\n", y)
	//fmt.Printf("Value of y in hexadecimal is %X\n", y)
	//fmt.Printf("Value of y in decimal is %v\n", y)

	//x := 455
	//var p *int
	//p = &x
	//
	//fmt.Println("Value stored in X: ", x)
	//fmt.Println("Address stored in X: ", &x)
	//fmt.Println("Address stored in P: ", p)
	//fmt.Println("Value stored in P: ", *p)

	//var pntr *int
	//fmt.Println("S = ", pntr)

	//TODO: passing pointer to the function
	x := 400
	fmt.Printf("x = %d\n", x)

	ptf(&x)
	fmt.Printf("x = %d\n", x)

}

func ptf(a *int) {
	*a = 748

}
