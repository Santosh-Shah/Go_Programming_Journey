package main

import "fmt"

func main() {
	//Pi = 41    not allowed to assign
	fmt.Println("Value of Pi:", Pi)
	fmt.Println("Name:", name)

	anotherName = "Roshan"
	fmt.Println("Another Name:", anotherName)
}

var anotherName = "rohit"

const Pi int = 45
const name string = "Samir Shah"
