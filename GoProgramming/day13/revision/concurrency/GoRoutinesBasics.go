package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello Vedas College")
		time.Sleep(1 * time.Second)
	}
}

func sayHii() {
	fmt.Println("Hii Vedas College")
}

func main() {
	//go sayHello()
	//fmt.Println("Hello Nation Infotech College")
	//time.Sleep(5 * time.Second)
	//go sayHii()
	//time.Sleep(1 * time.Second)

	go sayHello()
	fmt.Println("Hello from main function")
	time.Sleep(5 * time.Second)
}
