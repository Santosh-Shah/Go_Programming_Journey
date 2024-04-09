package main

import (
	"bytes"
	"fmt"
)

func main() {
	// slice with composite literal
	//sa := []int{1, 2, 3, 4, 5}
	//fmt.Println(sa)
	//
	//slice := []string{"apple", "banana", "cherry"}
	//fmt.Println("Slice: ", slice)
	//
	//sortSlice := []int{1, 32, 4, 56, 11}
	//sort.Ints(sortSlice)
	//fmt.Println("Sorted: ", sortSlice)
	// check if slice is sorted or not
	//isSliceSorted := sort.IntsAreSorted(sortSlice)
	//fmt.Println("is Sorted: ", isSliceSorted)
	//
	//slice_1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's', 'f',
	//	'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}
	//slice_2 := []byte{'*', '*', 'A', 'p', 'p', 'l', 'e', '^', '^'}
	//slice_3 := []byte{'%', 'g', 'e', 'e', 'k', 's', '%'}

	// Displaying slices
	//fmt.Println("Original Slice:")
	//fmt.Printf("Slice 1: %s", slice_1)
	//fmt.Printf("\nSlice 2: %s", slice_2)
	//fmt.Printf("\nSlice 3: %s", slice_3)

	// Trimming specified leading
	// and trailing Unicodes points
	// from the given slice of bytes
	// Using Trim function
	//res1 := bytes.Trim(slice_1, "!#")
	//res2 := bytes.Trim(slice_2, "*^")
	//res3 := bytes.Trim(slice_3, "@")

	// Display the results
	//fmt.Printf("New Slice:\n")
	//fmt.Printf("\nSlice 1: %s", res1)
	//fmt.Printf("\nSlice 2: %s", res2)
	//fmt.Printf("\nSlice 3: %s", res3)

	// Creating and trimming
	// the slice of bytes
	// Using Trim function
	//res1 := bytes.Trim([]byte("****Welcome to GeeksforGeeks****"), "*")
	//res2 := bytes.Trim([]byte("!!!!Learning how to trim a slice of bytes@@@@"), "!@")
	//res3 := bytes.Trim([]byte("^^Geek&&"), "$")

	// Display the results
	//fmt.Printf("Final Slice:\n")
	//fmt.Printf("\nSlice 1: %s", res1)
	//fmt.Printf("\nSlice 2: %s", res2)
	//fmt.Printf("\nSlice 3: %s", res3)

	//byteSlice := []byte("  Hello, World!  ")
	//trimmed := bytes.TrimSpace(byteSlice)
	//fmt.Println(string(trimmed)) // "Hello, World!"

	//TODO: split a slice of bytes
	// Creating and initializing
	// the slice of bytes
	// Using shorthand declaration
	//slice_1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's',
	//	'f', 'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}
	//
	//slice_2 := []byte{'A', 'p', 'p', 'l', 'e'}
	//
	//slice_3 := []byte{'%', 'g', '%', 'e', '%', 'e',
	//	'%', 'k', '%', 's', '%'}

	// Displaying slices
	//fmt.Println("Original Slice:")
	//fmt.Printf("Slice 1: %s", slice_1)
	//fmt.Printf("\nSlice 2: %s", slice_2)
	//fmt.Printf("\nSlice 3: %s", slice_3)

	// Splitting the slice of bytes
	// Using Split function
	//res1 := bytes.Split(slice_1, []byte("eek"))
	//res2 := bytes.Split(slice_2, []byte(""))
	//res3 := bytes.Split(slice_3, []byte("%"))

	// Display the results
	//fmt.Printf("\n\nAfter splitting:")
	//fmt.Printf("\nSlice 1: %s", res1)
	//fmt.Printf("\nSlice 2: %s", res2)
	//fmt.Printf("\nSlice 3: %s", res3)

	// Creating and Splitting
	// the slice of bytes
	// Using Split function
	res1 := bytes.Split([]byte("****Welcome, to, GeeksforGeeks****"),
		[]byte("Welcome"))

	res2 := bytes.Split([]byte("Learning x how x to x trim x a x slice of bytes"),
		[]byte("x"))

	res3 := bytes.Split([]byte("GeeksforGeeks, Geek"), []byte(""))

	res4 := bytes.Split([]byte(""), []byte(","))

	// Display the results
	fmt.Printf("Final Value:\n")
	fmt.Printf("\nSlice 1: %s", res1)
	fmt.Printf("\nSlice 2: %s", res2)
	fmt.Printf("\nSlice 3: %s", res3)
	fmt.Printf("\nSlice 4: %s", res4)
}
