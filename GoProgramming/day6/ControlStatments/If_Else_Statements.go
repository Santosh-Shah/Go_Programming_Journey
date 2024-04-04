package main

import "fmt"

func main() {

	//TODO: If statements
	//v := 400
	//if v < 1000 {
	//	fmt.Println("Value of V= ", v, " is Less than 1000")
	//}
	//
	//fmt.Println("Value of V= ", v)

	//TODO: if else statements
	//v := 5000
	//if v < 1000 {
	//	fmt.Println("Value of V= ", v, " is Less than 1000")
	//} else {
	//	fmt.Println("Value of V= ", v, " is not less than 1000")
	//}

	//TODO: nested if statements
	//v := 900
	//if v < 1000 {
	//	fmt.Println("Value of V= ", v, " is Less than 1000")
	//	if v > 500 {
	//		if v == 900 {
	//			fmt.Println("Is value is: ", v)
	//		}
	//	}
	//}

	//TODO: if else...if ladder
	v := 500
	if v == 100 {
		fmt.Println("It is equal to 100: ", v)
	} else if v == 500 {
		fmt.Println("It is equal to 500: ", v)
	} else if v == 900 {
		fmt.Println("It is equal to 900: ", v)
	} else if v == 100 {
		fmt.Println("It is equal to 1000: ", v)
	} else {
		fmt.Println("It is not among of them")
	}

}
