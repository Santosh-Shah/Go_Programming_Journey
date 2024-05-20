package main

import "fmt"

// Define a Rectangle struct
type Rectangle struct {
	width, height int
}

type Circle struct {
	radius float64
}

func (rec Rectangle) area() int {
	return rec.width * rec.height
}

func (cir Circle) area() float64 {
	return 3.14 * cir.radius * cir.radius
}

func main() {
	// create an instance of Rectangle
	rect := Rectangle{width: 45, height: 10}
	cir := Circle{6}

	// Call the area method on the rect instance
	fmt.Println("Area of Rectangle:", rect.area())

	//TODO: Methods with Same Name
	fmt.Println("Area of Circle:", cir.area())
}
