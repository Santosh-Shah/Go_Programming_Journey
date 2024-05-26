package main

import "fmt"

func main() {
	var str string
	var char rune
	fmt.Println("Enter a string: ")
	fmt.Scanln(&str)

	fmt.Println()

	fmt.Println("Enter a character: ")
	fmt.Scanf("%c", &char)

	count := countCharacter(str, char)
	fmt.Println(rune(char), "occurrence is of: ", count, " times")

}

func countCharacter(str string, char rune) int {
	count := 0
	for _, charValue := range str {
		if charValue == char {
			count++
		}
	}

	return count
}
