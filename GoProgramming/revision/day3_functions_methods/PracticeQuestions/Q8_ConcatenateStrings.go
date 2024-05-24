package main

import (
	"fmt"
)

func main() {
	var str1, str2 string

	fmt.Print("Enter first string: ")
	fmt.Scanf("%s\n", &str1)

	fmt.Print("Enter second string: ")
	fmt.Scanf("%s\n", &str2)

	output := printConcatenateStrings(str1, str2)
	fmt.Println("Concatenated String: ", output)
}

func printConcatenateStrings(str1, str2 string) string {
	return str1 + str2
}
