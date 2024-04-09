package main

import (
	"fmt"
	"strings"
)

func main() {
	//str1 := "hello"
	//str2 := "world"
	//str3 := "Hello"
	//str4 := "world"

	//fmt.Println(str1 == str2)
	//fmt.Println(str1 == str3)
	//fmt.Println(str1 == str4)
	//fmt.Println(str2 == str4)
	//
	//fmt.Println("-------------------------------")
	//fmt.Println(str1 != str2)

	//fmt.Println("Slice" > "slice")

	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))

	fmt.Println(strings.Compare("a", "c"))
}
