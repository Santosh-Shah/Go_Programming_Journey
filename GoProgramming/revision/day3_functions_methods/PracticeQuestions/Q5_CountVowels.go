package main

import "fmt"

func main() {
	var input string

	fmt.Println("Enter a string: ")
	fmt.Scanf("%s", &input)
	output := countVowels(input)
	fmt.Println("Total numbers of vowels: ", output)
}

func countVowels(input string) int {
	vowels := "aeiouAEIOU"
	count := 0

	for _, char := range input {
		if checkVowel(vowels, char) {
			count++
		}
	}

	return count
}

func checkVowel(vowels string, char rune) bool {
	for _, vowel := range vowels {
		if vowel == char {
			return true
		}
	}
	return false
}
