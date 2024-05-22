package main

import (
	"fmt"
)

func sendOnly(ch chan<- string) {
	ch <- "Send Only"
}

func receiveOnly(ch <-chan string) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan string)

	go sendOnly(ch)
	//go receiveOnly(ch)
	fmt.Println(<-ch)

	// GIve some time fo goroutine to complete
	fmt.Scanln()
}
