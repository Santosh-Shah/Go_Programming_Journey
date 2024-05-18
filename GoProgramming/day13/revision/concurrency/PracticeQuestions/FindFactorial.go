package main

import "fmt"

func factorial(n int, ch chan<- int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	ch <- result
}

func main() {
	n := 5
	ch := make(chan int)

	go factorial(n, ch)
	result := <-ch
	fmt.Println("Result of factorial: ", result)
}
