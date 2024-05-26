package main

import "fmt"

func main() {
	str := "aba"

	output := isPalindrome(str)
	if output {
		fmt.Println("Yes it's Palindrome")
	} else {
		fmt.Println("No it's Palindrome")
	}
}

func isPalindrome(str string) bool {
	if len(str) == 0 {
		return true
	}

	i := 0
	j := len(str) - 1

	for i < j {
		if str[i] == str[j] {
			return true
		}
		i++
		j--
	}
	return false
}
