package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("----------Main function-----------")
	go printHello()
	time.Sleep(2 * time.Second)
	go printCollege()
	//time.Sleep(5 * time.Second)
	fmt.Println("--------------End of main function------------")

}

func printHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello world")
		//time.Sleep(1 * time.Second)
	}
}

func printCollege() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println("College")
	//}

	fmt.Println("College")
}
