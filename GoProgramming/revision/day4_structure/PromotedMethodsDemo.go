package main

import "fmt"

type Address struct {
	Country, City string
}

type Person struct {
	Name, CollegeName string
	Address
}

func (adrs Address) fullAddress() string {
	return adrs.Country + "-" + adrs.City
}

func main() {
	p := Person{
		Name:        "Santosh Shah",
		CollegeName: "Vedas College",
		Address: Address{
			Country: "Nepal",
			City:    "Lalitpur",
		},
	}

	fmt.Println("Full-Address:", p.fullAddress())
}
