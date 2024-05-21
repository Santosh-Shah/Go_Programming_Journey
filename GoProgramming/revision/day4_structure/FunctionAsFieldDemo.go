package main

import "fmt"

type SumOfTwoNumbers struct {
	add func(a, b int) int
}

func main() {
	sum := SumOfTwoNumbers{
		add: func(a, b int) int {
			return a + b
		},
	}

	fmt.Println(sum.add(45, 10))
}
