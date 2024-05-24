package main

import "fmt"

type Rectangle struct {
	length, width float64
}

func (rec Rectangle) area() float64 {
	output := rec.length * rec.width
	return output
}

func (rec Rectangle) perimeter() float64 {
	output := 2 * (rec.length + rec.width)
	return output
}

func main() {
	var length, width float64

	fmt.Printf("Enter length: ")
	fmt.Scan(&length)

	fmt.Printf("Enter width: ")
	fmt.Scan(&width)

	rec := Rectangle{length, width}

	fmt.Println("Perimeter of rectangle:", rec.perimeter())
	fmt.Println("Area of rectangle:", rec.area())

}
