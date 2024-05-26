package main

import (
	"fmt"
)

func extractSubstrings(str string, length int) []string {
	var substrings []string
	for i := 0; i <= len(str)-length; i++ {
		substrings = append(substrings, str[i:i+length])
	}
	return substrings
}

func main() {
	var inputStr string
	var length int
	fmt.Print("Enter a string: ")
	fmt.Scanln(&inputStr)
	fmt.Print("Enter the length of substrings: ")
	fmt.Scanln(&length)

	substrings := extractSubstrings(inputStr, length)
	fmt.Println("Substrings of length", length, ":", substrings)
}
