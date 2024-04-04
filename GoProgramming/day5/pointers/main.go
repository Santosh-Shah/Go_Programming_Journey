package main

import "fmt"

func main() {
	//var x int = 100
	//fmt.Println("value of x: ", x)
	//
	//var y = &x
	////var val = y
	//fmt.Println("Address of x: ", y)
	//fmt.Println("Value of x: ", *y)
	//
	//var z = &y
	//fmt.Println("Address of y: ", z)

	//TODO: if we do not initialize pointer then by default it <nil>
	//var pntr *int
	//fmt.Println("pointer: ", pntr)

	//TODO: shorthand syntax
	y := 500
	p := &y
	fmt.Println("Value of y: ", y)
	fmt.Println("Address of Y: ", p)
	fmt.Println("Address of Y: ", &y)
	fmt.Println("Value stored in P: ", *p)

	// now change the value of y
	*p = 1000
	fmt.Println("New Value of y: ", y)

}
