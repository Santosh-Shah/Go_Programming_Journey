package main

import "fmt"

type Employee struct {
	UserId   int
	Fullname string
	Username string
	Password string
	Phone    string
	Email    string
}

var employees []*Employee

func main() {
	// Creating instances of Employee struct
	emp1 := &Employee{
		UserId:   1,
		Fullname: "Ravi Gurung",
		Username: "ravi679",
		Password: "123456",
		Phone:    "9811243552",
		Email:    "ravi679@gmail.com",
	}

	emp2 := &Employee{
		UserId:   1,
		Fullname: "Prem Gurung",
		Username: "prem679",
		Password: "123456",
		Phone:    "9849697391",
		Email:    "prem679@gmail.com",
	}

	employees = append(employees, emp1, emp2)

	fmt.Println("All Employees: ")
	for value, emp := range employees {
		fmt.Print(value, " ")
		fmt.Println("employee phone no: %s\n", emp.Phone)
	}
}
