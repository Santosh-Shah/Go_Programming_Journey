package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
}

func newPerson(name, gender string, age int) Person {
	return Person{name, age, gender}
}

func printPerson(person Person) {
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Gender:", person.Gender)
}

func main() {
	var name, gender string
	var age int
	print("Enter your name: ")
	fmt.Scan(&name)
	print("Enter your age: ")
	fmt.Scan(&age)
	print("Enter your gender: ")
	fmt.Scan(&gender)

	person := newPerson(name, gender, age)
	printPerson(person)
}
