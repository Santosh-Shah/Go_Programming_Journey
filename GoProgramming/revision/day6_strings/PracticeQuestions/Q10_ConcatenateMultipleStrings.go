package main

import (
	"fmt"
	"strings"
)

func concatenateStrings(strs ...string) string {
	return strings.Join(strs, "")
}

func main() {
	var inputStrings []string
	var inputStr string
	var n int

	fmt.Print("Enter the number of strings: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter string %d: ", i+1)
		fmt.Scanln(&inputStr)
		inputStrings = append(inputStrings, inputStr)
	}

	concatenatedString := concatenateStrings(inputStrings...)
	fmt.Println("Concatenated string:", concatenatedString)
}
