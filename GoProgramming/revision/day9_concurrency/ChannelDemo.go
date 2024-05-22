package main

import "fmt"

func sendMsg(ch chan string) {
	ch <- "Hello from goroutine hello hello"
}

func main() {
	chh := make(chan string)
	go sendMsg(chh)

	// Receiving message from channel
	msg := <-chh
	fmt.Println(msg)
}
