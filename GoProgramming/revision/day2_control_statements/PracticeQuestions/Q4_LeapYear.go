package main

import "fmt"

func findLeapYear(year int) {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println(year, ": is a leap year")
	} else {
		fmt.Println(year, ": is not a leap year")
	}
}

func main() {
	num := 2004
	findLeapYear(num)
}
