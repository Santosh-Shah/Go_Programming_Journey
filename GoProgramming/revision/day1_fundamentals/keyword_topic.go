package main

import "fmt"

func main() {
	var x int = 45 //'var is a keyword
	if x < 10 {    // if and else are also keywords
		fmt.Println("It is less than 10")
	} else {
		fmt.Println("It is either greater or equal to 10")
	}
}
