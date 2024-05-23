package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number to print *: ")
	fmt.Scan(&num)

	printStartPattern(num)
}

func printStartPattern(num int) {
	for i := 0; i < num; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
