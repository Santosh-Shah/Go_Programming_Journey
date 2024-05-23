package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number to calculate sum of divisible of 3 and 5: ")
	fmt.Scan(&num)

	fmt.Println(numberDivisibleOf3and5(num))
	output := sumOfDivisible(num)
	fmt.Println("The sum of divisible of 3 and 5 is", output)
}

func sumOfDivisible(num int) int {
	sum := 0

	for i := 1; i <= num; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func numberDivisibleOf3and5(num int) []int {
	slice := make([]int, 0)

	for i := 1; i <= num; i++ {
		if i%3 == 0 || i%5 == 0 {
			slice = append(slice, i)
		}
	}

	return slice
}
