package main

import (
	"fmt"
	"time"
)

func main() {
	go display("Welcome")

	display("Vedas College")

	//TODO: Anonymous Goroutine
	//fmt.Println("Welcome!! to Main function")
	//go func() {
	//	fmt.Println("Welcome!! to GeeksForGeeks")
	//}()
	//
	//time.Sleep(1 * time.Second)
	//fmt.Println("GoodBye!! to Main function")
}

func display(value string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(value)
	}
}
