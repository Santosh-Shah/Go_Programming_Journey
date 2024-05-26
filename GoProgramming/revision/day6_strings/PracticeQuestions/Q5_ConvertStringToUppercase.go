package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "vedas college"
	fmt.Println("Original String:", str)

	output := convertUpperCase(str)
	fmt.Println("Uppercase String:", output)
}

func convertUpperCase(str string) string {
	return strings.ToUpper(str)
}
