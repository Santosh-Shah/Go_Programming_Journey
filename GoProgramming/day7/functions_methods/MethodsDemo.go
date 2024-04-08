package main

import "fmt"

// Author structure
type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

// Method with a receiver
// of author type
func (a author) show() {
	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

// Type definition
type data int

// Defining a method with
// non-struct type receiver
func (d1 data) add(d2 data) data {
	return d1 * d2
}

/*
// if you try to run this code,
// then compiler will throw an error
func(d1 int)multiply(d2 int)int{
return d1 * d2
}
*/

// Main function
func main() {

	// Initializing the values
	// of the author structure
	//res := author{
	//	name:      "Sona",
	//	branch:    "CSE",
	//	particles: 203,
	//	salary:    34000,
	//}
	//
	//// Calling the method
	//res.show()

	value1 := data(20)
	value2 := data(20)
	res := value1.add(value2)
	fmt.Println("Final result: ", res)
}
