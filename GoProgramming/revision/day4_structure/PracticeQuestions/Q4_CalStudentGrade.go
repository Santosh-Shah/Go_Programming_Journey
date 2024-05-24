package main

import "fmt"

type StudentInfo struct {
	name   string
	rollNo int64
	marks  float64
}

func (student StudentInfo) calculateGrade() string {
	if student.marks >= 80 && student.marks <= 100 {
		return "A"
	} else if student.marks < 80 && student.marks >= 60 {
		return "B"
	} else if student.marks < 60 && student.marks >= 40 {
		return "c"
	} else {
		return "D"
	}
}

func printStudentDetails(std StudentInfo) {
	fmt.Println("Student Details:")
	fmt.Println("Student Name: ", std.name)
	fmt.Println("Student Roll No: ", std.rollNo)
	fmt.Println("Student Marks: ", std.marks)
	fmt.Println("Student Grade: ", std.calculateGrade())
}

func main() {

	var name string
	var marks float64
	var rollNo int64

	fmt.Print("Enter name: ")
	fmt.Scanf("%s\n", &name)

	fmt.Print("Enter roll no: ")
	fmt.Scanf("%d\n", &rollNo)

	fmt.Print("Enter marks: ")
	fmt.Scanf("%f\n", &marks)

	student := StudentInfo{name, rollNo, marks}
	printStudentDetails(student)

}
