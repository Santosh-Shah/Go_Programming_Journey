package main

import "fmt"

func main() {
	str := "Hello World"
	output := lengthOfString(&str)
	fmt.Println("Length of String:", output)
}

func lengthOfString(str *string) int {
	return len(*str)
}
