package main

import "fmt"

//func send(ch chan<- int) {
//	ch <- 42 // Send data into the channel
//}

//func receive(ch <-chan int) {
//	msg := <-ch // Receive data from the channel
//	fmt.Println("Received:", msg)
//}

func send(ch1 chan<- int) {
	ch1 <- 4444
}

func receive(ch1 <-chan int) {
	value := <-ch1
	fmt.Println("Address: ", ch1)
	fmt.Println("Value: ", value)

}

func main() {
	//ch := make(chan int)
	//go send(ch) // Send-only channel
	//receive(ch) // Receive-only channel

	ch1 := make(chan int)

	go send(ch1)
	receive(ch1)

}
