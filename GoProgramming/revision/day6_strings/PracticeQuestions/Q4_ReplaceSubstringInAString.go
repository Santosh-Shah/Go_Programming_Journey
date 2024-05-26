package main

import (
	"fmt"
	"strings"
)

func replaceSubstring(str, oldStr, newStr string) string {
	return strings.ReplaceAll(str, oldStr, newStr)
}

func main() {
	var inputStr, oldSub, newSub string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&inputStr)
	fmt.Print("Enter the substring to replace: ")
	fmt.Scanln(&oldSub)
	fmt.Print("Enter the new substring: ")
	fmt.Scanln(&newSub)

	modifiedStr := replaceSubstring(inputStr, oldSub, newSub)
	fmt.Println("Modified string:", modifiedStr)
}
