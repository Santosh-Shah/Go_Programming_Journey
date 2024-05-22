package main

//import (
//	"fmt"
//	"sync"
//)
//
//func sayHello() {
//	fmt.Println("Hello World")
//}
//
//func printMsg(msg string, wg *sync.WaitGroup) {
//	defer wg.Done()
//	fmt.Println(msg)
//}
//
//func main() {
//	//go sayHello()
//	//fmt.Println("main function")
//	//time.Sleep(1 * time.Second)
//
//	//TODO: multiple goroutines
//	var wg sync.WaitGroup
//	messages := []string{"hello", "world", "!"}
//	for _, msg := range messages {
//		wg.Add(1)
//		go printMsg(msg, &wg)
//	}
//
//	wg.Wait()
//
//}
