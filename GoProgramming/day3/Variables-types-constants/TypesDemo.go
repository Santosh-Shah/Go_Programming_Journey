package main

import "fmt"

func main() {

	//TODO: Unsigned Integer
	var uint8Var uint8 = 255
	var uint16Var uint16 = 65535
	var uint32Var uint32 = 4294967295
	var uint64Var uint64 = 18446744073709551615

	//TODO: Signed Integer
	var int8Var int8 = -128
	var int16Var int16 = -32768
	var int32Var int32 = -2147483648
	var int64Var int64 = -9223372036854775808

	//TODO: Floating numbers
	var float32Var float32 = 3.14159265358979323846264338327950288419716939937510582097494459
	var float64Var float64 = 3.14159265358979323846264338327950288419716939937510582097494459

	//TODO: Complex Numbers
	var complex64Var complex64 = 3 + 4i
	var complex128Var complex128 = 3 + 4i

	fmt.Println("Unsigned Integers numbers: ", uint8Var, uint16Var, uint32Var, uint64Var)
	fmt.Println("Signed Integers Numbers: ", int8Var, int16Var, int32Var, int64Var)
	fmt.Println("Floating numbers: ", float32Var, float64Var)
	fmt.Println("Complex Numbers: ", complex64Var, complex128Var)

	// no var style
	numberOfUser := 253
	name := "Hariom Shah"
	age := "23"
	isMarried := false

	fmt.Println(numberOfUser)
	fmt.Println(LoginToken)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
}

const LoginToken = "Vedas College"
