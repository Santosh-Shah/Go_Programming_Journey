package main

import "fmt"

func main() {
	// Shorthand declaration of array
	//arr := [4]string{"geek", "gfg", "geeks1522", "GeeksForGeeks"}
	//
	//fmt.Println("Elements of the array:")
	//fmt.Println(arr)

	// print the value of array using for loop
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}

	//TODO: multidimensional array
	//arr := [3][3]string{
	//	{"C#", "Python", "Java"},
	//	{"Golang", "Ruby", "C++"},
	//	{"Javascript", "Rust", "Spring Boot"},
	//}

	// accessing the values of the
	// array using for loop
	//fmt.Println("Elements of Arrays:")
	//for i := 0; i < len(arr); i++ {
	//	for j := 0; j < len(arr[i]); j++ {
	//		fmt.Print(arr[i][j], " ")
	//	}
	//	fmt.Println()
	//}

	//fmt.Println("------------------------")

	//TODO: another way to implement array
	//var arr1 [2][2]int
	//arr1[0][0] = 1
	//arr1[0][1] = 2
	//arr1[1][0] = 3
	//arr1[1][1] = 4

	//fmt.Println("Elements of array2: ")
	//for i := 0; i < len(arr1); i++ {
	//	for j := 0; j < len(arr1[i]); j++ {
	//		fmt.Print(arr1[i][j], " ")
	//	}
	//	fmt.Println()
	//}

	//TODO: Creating a array of int type
	//var myarr [2]int
	//myarr[0] = 1
	//fmt.Println(myarr)

	//TODO: len() method
	//arr1 := [3]int{1, 2, 3}
	//arr2 := [4]int{12, 34, 56, 77}
	//fmt.Println("Length of arr1: ", len(arr1))
	//fmt.Println("Length of arr2: ", len(arr2))

	//TODO: using ellipsis (...)
	//myarray := [...]string{
	//	"GFG", "GfG", "geeks", "GeeksForGeeks", "Geek"}
	//
	//fmt.Println("Elements of the array: ", myarray)
	//fmt.Println("Length of the array: ", len(myarray))

	//TODO: creating an array whose size is represented by the ellipsis
	//myarray := [...]int{
	//	12, 43, 56, 78, 22, 44}
	//
	//for i := 0; i < len(myarray); i++ {
	//	fmt.Printf("%d\n", myarray[i])
	//}

	//myarray := [...]int{12, 33, 44, 55, 66, 77}
	//fmt.Println("Original array(Before): ", myarray)
	//
	//newArray := myarray
	//fmt.Println("New array:(before): ", newArray)
	//newArray[0] = 500
	//
	//fmt.Println("New array(After): ", newArray)
	//
	//fmt.Println("Original array(After): ", myarray)

	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}
	arr3 := [3]int{1, 20, 3}

	// Comparing arrays using == operator
	fmt.Println(arr1 == arr2)
	fmt.Println(arr2 == arr3)
	fmt.Println(arr1 == arr3)
}
