package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "Msg from channel one"
	}()

	go func() {
		time.Sleep(10 * time.Second)
		ch2 <- "Msg from channel two"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Message Received: ", msg1)
	case msg2 := <-ch2:
		fmt.Println("Message Received: ", msg2)
	}

	//// Receive from ch1
	//msg1 := <-ch1
	//fmt.Println("Received:", msg1)
	//
	//// Receive from ch2
	//msg2 := <-ch2
	//fmt.Println("Received:", msg2)
}
