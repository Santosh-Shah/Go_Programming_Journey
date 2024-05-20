package main

import "fmt"

func main() {
	printName()
	defer clearUp()
	fmt.Println("This is from Main function:")
}

func clearUp() {
	fmt.Println("This is from clearUp function:")
}

func printName() {
	fmt.Println("My college name is Vedas College:")
}
