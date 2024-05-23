package main

import "fmt"

func fibonacci(num int) []int {

	output := make([]int, num)
	if num <= 0 || num <= 1 {
		output[0] = -1
		return output
	}

	output[0] = 0
	output[1] = 1
	for i := 2; i < num; i++ {
		output[i] = output[i-1] + output[i-2]
	}
	return output
}

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)

	output := fibonacci(num)
	fmt.Println(output)
}
