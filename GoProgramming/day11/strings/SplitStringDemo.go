package main

import (
	"fmt"
	"strings"
)

func main() {
	//	// Creating and initializing the strings
	//	str1 := "Welcome, to the, online portal, of GeeksforGeeks"
	//	str2 := "My dog name is Dollar"
	//	str3 := "I like to play Ludo"
	//
	//	// Displaying strings
	//	fmt.Println("String 1: ", str1)
	//	fmt.Println("String 2: ", str2)
	//	fmt.Println("String 3: ", str3)
	//
	//	// Splitting the given strings
	//	// Using Split() function
	//	res1 := strings.Split(str1, ",")
	//	res2 := strings.Split(str2, "")
	//	res3 := strings.Split(str3, "!")
	//	res4 := strings.Split("", "GeeksforGeeks, geeks")
	//
	//	// Displaying the result
	//
	//	fmt.Println("\nResult 1: ", res1)
	//	fmt.Println("Result 2: ", res2)
	//	fmt.Println("Result 3: ", res3)
	//	fmt.Println("Result 4: ", res4)
	//}

	// Creating and initializing the strings
	//str1 := "Welcome, to the, online portal, of GeeksforGeeks"
	//str2 := "My dog name is Dollar"
	//str3 := "I like to play Ludo"
	//
	//// Displaying strings
	//fmt.Println("String 1: ", str1)
	//fmt.Println("String 2: ", str2)
	//fmt.Println("String 3: ", str3)
	//
	//// Splitting the given strings
	//// Using SplitAfter() function
	//res1 := strings.SplitAfter(str1, ",")
	//res2 := strings.SplitAfter(str2, "")
	//res3 := strings.SplitAfter(str3, "!")
	//res4 := strings.SplitAfter("", "GeeksforGeeks, geeks")
	//
	//// Displaying the result
	//fmt.Println("\nResult 1: ", res1)
	//fmt.Println("Result 2: ", res2)
	//fmt.Println("Result 3: ", res3)
	//fmt.Println("Result 4: ", res4)

	// Creating and initializing the strings
	str1 := "Welcome, to the, online portal, of GeeksforGeeks"
	str2 := "My dog name is Dollar"
	str3 := "I like to play Ludo"

	// Displaying strings
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)
	fmt.Println("String 3: ", str3)

	// Splitting the given strings
	// Using SplitAfterN() function
	res1 := strings.SplitAfterN(str1, ",", 2)
	res2 := strings.SplitAfterN(str2, "", 4)
	res3 := strings.SplitAfterN(str3, "!", 1)
	res4 := strings.SplitAfterN("", "GeeksforGeeks, geeks", 3)

	// Displaying the result
	fmt.Println("\nResult 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4: ", res4)
}
