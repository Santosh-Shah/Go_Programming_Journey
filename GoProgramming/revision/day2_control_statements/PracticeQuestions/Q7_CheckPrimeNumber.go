package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number to check the prime number: ")
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Println(num, " :number is prime")
	} else {
		fmt.Println(num, " :number is not prime")
	}
}

func isPrime(num int) bool {

	// if the number is less than or equal to 1 then it is not prime
	if num <= 1 {
		return false
	}

	//  it will check if product of each iteration with num
	// if it divisible then not prime
	// it also reduce the time of checking
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
