package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
		//time.Sleep(5 * time.Second)
	}

	wg.Wait()
	fmt.Println("End of main function")
	//time.Sleep(2 * time.Second)
}

func worker(count int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker start: %d\n", count)
	fmt.Printf("worker end: %d\n", count)
}
