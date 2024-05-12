package main

import "fmt"

func main() {
	//ch := make(chan string) // Create a channel
	//
	//go func() {
	//	ch <- "Hello" // Send data into the channel
	//}()
	//
	//msg := <-ch      // Receive data from the channel
	//fmt.Println(msg) // Prints: Hello

	ch := make(chan int)

	go func() {
		ch <- 45
	}()

	value := <-ch
	fmt.Println("value", value)
}
