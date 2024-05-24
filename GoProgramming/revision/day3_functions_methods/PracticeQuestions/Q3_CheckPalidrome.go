package main

import "fmt"

func main() {
	var num string
	fmt.Println("Enter number to check palindrome: ")
	fmt.Scanln(&num)

	output := isPalindrome(num)
	if output == true {
		fmt.Println(num, "is a palindrome")
	} else {
		fmt.Println(num, "is not a palindrome")
	}

}

func isPalindrome(num string) bool {
	if len(num) <= 1 {
		return true
	}

	len := len(num)
	for i := 0; i < len/2; i++ {
		if num[i] != num[len-i-1] {
			return false
		}
	}

	return true
}
