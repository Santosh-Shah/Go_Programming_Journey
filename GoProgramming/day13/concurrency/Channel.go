package main

import "fmt"

func main() {
	// Creating a channel
	// Using var keyword
	var mychannel chan int
	fmt.Println("Value of the channel: ", mychannel)
	fmt.Printf("Type of the channel: %T\n ", mychannel)

	// creating a channel using make() function
	mychannel1 := make(chan int)
	fmt.Println("Value of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T\n ", mychannel1)

}
