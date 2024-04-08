package main

import (
	"fmt"
)

func main() {
	var a Address
	fmt.Println(a)

	// Declaring and initializing it
	//a1 := Address{"Hariom Shah", "Kalaiya", 4521}
	//fmt.Println("Address:", a1)
	//
	//a2 := Address{name: "Santosh Shah", city: "Kathmandu", pinCode: 1111}
	//fmt.Println("Address2:", a2)
	//
	//a3 := Address{name: "Rohit"}
	//fmt.Println("Address3:", a3)

	//TODO: lets access the field of struct
	c := Address{name: "Ramu", city: "Birgunj", pinCode: 5555}
	fmt.Println("Person1 Name: ", c.name)
	fmt.Println("Person1 City: ", c.city)

	// Assigning a new value
	// to a struct field
	c.pinCode = 4444
	fmt.Println("Person2 PinCode: ", c.pinCode)
	fmt.Println("-------------------------------------")

	emp1 := &Employee{"Shyam", "Tamang", 55, 1000}

	// one way to access
	fmt.Println("First Name: ", (*emp1).firstName)
	fmt.Println("Age: ", (*emp1).age)

	// other way to access
	fmt.Println("First Name: ", emp1.firstName)
	fmt.Println("Age: ", emp1.age)

	// now explore other use of structure
	employee := Employee{
		firstName: "Raju",
		lastName:  "Yadav",
		age:       25,
		salary:    25000,
	}
	fmt.Println(employee)

}

// Employee struct
type Employee struct {
	firstName, lastName string
	age, salary         int
}

// creating a struct type
type Address struct {
	name    string
	city    string
	pinCode int
}
