package main

import (
	"fmt"
	"strings"
)

func main() {
	//str := "Hello everyone how are you!"
	//fmt.Println("String values: ", str)

	//TODO: comparison of string
	//str1 := "Ram"
	//str2 := "ram"
	//str3 := "shyam"
	//str4 := "Ram"
	//fmt.Println(str1 == str2)
	//fmt.Println(str1 == str3)
	//fmt.Println(str1 == str4)
	//
	//fmt.Println()
	//
	//fmt.Println(strings.Compare(str1, str2))
	//fmt.Println(strings.Compare(str1, str3))
	//fmt.Println(strings.Compare(str1, str4))

	//TODO: Ways to concatenate two strings
	//str1 := "Hello"
	//str2 := "World"
	//result := str1 + ", " + str2
	//fmt.Println("Concatenated String: ", result)
	//fmt.Println("Another ways:")
	//
	//result1 := fmt.Sprintf("%s, %s", str1, str2)
	//fmt.Println("Concatenated String: ", result1)
	//
	//fmt.Println("Another ways:")
	//result2 := strings.Join([]string{str1, str2}, ", ")
	//fmt.Println("Concatenated String: ", result2)

	//TODO: Trimming a string
	//TODO: TrimSpace()
	//str := "   Hello World!  "
	//fmt.Println(str)
	//afterSpaceTrimmed := strings.TrimSpace(str)
	//fmt.Println(afterSpaceTrimmed)

	//TODO: TrimPrefix() and TrimSuffix()
	//str := "!!!Hello World!!!"
	//trimmedPrefix := strings.TrimPrefix(str, "!!")
	//trimmedSuffix := strings.TrimSuffix(str, "!!!")
	//fmt.Println("Trimmed Prefix:", trimmedPrefix)
	//fmt.Println("Trimmed Suffix:", trimmedSuffix)
	//fmt.Println(str)

	//TODO: Splitting  a string
	str := "Hello World,Go,Programming"
	subStrings := strings.Split(str, "o")
	fmt.Println("Substrings: ", subStrings)

	//TODO: check if the given character is present or not
	fmt.Println("Contains 'Hello':", strings.Contains(str, "r"))

	//TODO: Repeat a string for specific number of time
	str1 := "Santosh "
	repeatedStr := strings.Repeat(str1, 5)
	fmt.Println("Repeated String:", repeatedStr)

	//TODO: finding the index value of a specific string
	index := strings.Index(str1, "s")
	fmt.Println(index)

	//TODO: Counting the number of Repeated Characters
	str2 := "Hello World, Lovely world"
	count := strings.Count(str2, "o")
	fmt.Println(count)

}
