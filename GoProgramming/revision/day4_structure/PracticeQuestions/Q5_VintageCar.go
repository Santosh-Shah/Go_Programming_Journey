package main

import (
	"fmt"
	"time"
)

type Car struct {
	make, model string
	year        int
}

func isVintageCar(c Car) bool {
	if (time.Now().Year() - c.year) >= 25 {
		return true
	}
	return false
}

func main() {
	var make, model string
	var year int

	fmt.Print("Enter car make: ")
	fmt.Scanf("%s\n", &make)

	fmt.Print("Enter car model:")
	fmt.Scanf("%s\n", &model)

	fmt.Print("Enter car year: ")
	fmt.Scanf("%d\n", &year)

	vintageCar := Car{make, model, year}
	if isVintageCar(vintageCar) {
		fmt.Println(make, ":is Vintage Car")
	} else {
		fmt.Print(make, ":is not a Vintage Car")
	}

}
