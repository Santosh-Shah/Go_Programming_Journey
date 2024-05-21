package main

import "fmt"

func main() {
	//TODO: Define am anonymous struct
	P := struct {
		name    string
		address string
		age     int
	}{
		name:    "Hariom Shah",
		address: "Malahani",
		age:     20,
	}

	fmt.Println(P)
	fmt.Println(P.name)
	fmt.Println(P.age)
	fmt.Println(P.address)
}
