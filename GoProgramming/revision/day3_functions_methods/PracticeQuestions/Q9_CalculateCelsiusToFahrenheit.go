package main

import "fmt"

func main() {
	var celsius float64
	fmt.Println("Enter your Celsius: ")
	fmt.Scanf("%f", &celsius)

	output := calculateFaht(celsius)
	fmt.Println("Your fahrenheit is ", output)
}

func calculateFaht(calsius float64) float64 {
	return calsius*(9/5) + 32
}
