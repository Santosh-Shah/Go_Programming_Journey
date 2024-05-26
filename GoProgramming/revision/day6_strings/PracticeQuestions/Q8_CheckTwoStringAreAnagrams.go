package main

import (
	"fmt"
	"sort"
	"strings"
)

func areAnagrams(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	// Sort the characters in both strings
	sortedStr1 := sortString(str1)
	sortedStr2 := sortString(str2)

	// Compare the sorted strings
	return sortedStr1 == sortedStr2
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	var str1, str2 string
	fmt.Print("Enter the first string: ")
	fmt.Scanln(&str1)
	fmt.Print("Enter the second string: ")
	fmt.Scanln(&str2)

	if areAnagrams(str1, str2) {
		fmt.Println("The strings are anagrams.")
	} else {
		fmt.Println("The strings are not anagrams.")
	}
}
