package main

import "fmt"

func main() {
	emp := Employee{"ABC", 12345}
	pts := &emp

	fmt.Println(pts)

	// updating the value
	//pts.name = "hello"
	//fmt.Println(pts.name)
	//fmt.Println((*pts).name)
}

type Employee struct {
	name  string
	empid int
}
