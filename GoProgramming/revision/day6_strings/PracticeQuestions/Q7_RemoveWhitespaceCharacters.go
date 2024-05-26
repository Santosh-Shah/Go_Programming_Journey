package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "my   name   is khan"
	fmt.Println("Original string: ", str)

	output := removeWhiteSpace(str)
	fmt.Println("After removeWhiteSpace: ", output)
}

func removeWhiteSpace(str string) string {
	return strings.ReplaceAll(str, "  ", "")
}
