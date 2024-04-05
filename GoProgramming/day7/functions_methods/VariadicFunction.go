package main

import (
	"fmt"
	"strings"
)

func main() {

	//TODO: Variadic function using string
	// it will print nothing
	//fmt.Println(joinString())
	//
	//// with value
	//fmt.Println(joinString("Santosh", "Hariom"))
	//fmt.Println(joinString("Raju", "Ravi", "Rohit"))
	//fmt.Println(joinString("Ram", "Shyam", "Mohit", "Hari"))

	//TODO: using Integer value
	fmt.Println(variadicSum(1, 5, 6))
	fmt.Println(variadicSum(1, 50, 60, 70))
	fmt.Println(variadicSum(1, 5))

}

func variadicSum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func joinString(elements ...string) string {
	return strings.Join(elements, "-")
}
