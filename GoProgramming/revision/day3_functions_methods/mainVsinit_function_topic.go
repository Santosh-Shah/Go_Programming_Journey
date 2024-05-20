package main

import "fmt"

func init() {
	fmt.Println("This is from init function 1:")
}

func main() {
	fmt.Println("This is from Main function:")
}

func init() {
	fmt.Println("This is from init function 2:")
}
