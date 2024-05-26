package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

// For Rectangle
type Rectangle struct {
	length, width float64
}

func (rec Rectangle) Area() float64 {
	return rec.length * rec.width
}

// for Circle
type Circle struct {
	radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func main() {
	rec := Rectangle{10.5, 6.3}
	circle := Circle{8.5}

	shapes := []Shape{rec, circle}

	// just for understanding
	fmt.Println(len(shapes))
	fmt.Println(shapes)

	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}
