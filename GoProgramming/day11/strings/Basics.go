package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//name := "Santosh Shah"
	//collegeName := "Vedas College"
	//fmt.Println("Name: ", name)
	//fmt.Println("College Name: ", collegeName)

	//value1 := "Welcome to geeksforgeeks"
	//value2 := "Welcome to v\nedas college"
	//value3 := `Hello everyone`
	//value4 := `Hello\n vedas college`
	//
	//fmt.Println("String1: ", value1)
	//fmt.Println("String2: ", value2)
	//fmt.Println("String3: ", value3)
	//fmt.Println("String4: ", value4)

	// Creating and initializing a string
	// using shorthand declaration
	//mystr := "Welcome to GeeksforGeeks"
	//fmt.Println("String:", mystr)

	//if you trying to change
	//      the value of the string
	//      then the compiler will
	//      throw an error, i.e,
	//     cannot assign to mystr[1]

	//mystr[1] = 'G'
	//fmt.Println("String:", mystr)

	// String as a range in the for loop
	//for index, s := range "GeeksForGeeKs" {
	//	fmt.Printf("The index number of %c is %d\n", s, index)
	//}

	// Creating and initializing a string

	//TODO: creating the string using slice
	// Creating and initializing a slice of byte
	myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}

	// Creating a string from the slice
	mystring1 := string(myslice1)

	// Displaying the string
	fmt.Println("String 1: ", mystring1)

	// Creating and initializing a slice of rune
	myslice2 := []rune{0x0047, 0x0065, 0x0065,
		0x006b, 0x0073}

	// Creating a string from the slice
	mystring2 := string(myslice2)

	// Displaying the string
	fmt.Println("String 2: ", mystring2)

	//TODO: finding the length of the string
	// Creating and initializing a string
	// using shorthand declaration
	mystr := "Welcome to GeeksforGeeks ??????"

	// Finding the length of the string
	// Using len() function
	length1 := len(mystr)

	// Using RuneCountInString() function
	length2 := utf8.RuneCountInString(mystr)

	// Displaying the length of the string
	fmt.Println("string:", mystr)
	fmt.Println("Length 1:", length1)
	fmt.Println("Length 2:", length2)
}
