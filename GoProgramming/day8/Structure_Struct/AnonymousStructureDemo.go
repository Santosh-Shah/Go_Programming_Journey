package main

import "fmt"

func main() {
	//student := struct {
	//	name, subject string
	//	age           int
	//}{
	//	name:    "Hariom Shah",
	//	subject: "Commerce",
	//	age:     25,
	//}
	//
	//fmt.Println(student)

	//TODO: now create anonymous fields
	type student struct {
		int
		string
		float64
	}

	value := student{45, "Hello", 45.3}
	fmt.Println(value.int, "\n", value.string, "\n", value.float64)
}
