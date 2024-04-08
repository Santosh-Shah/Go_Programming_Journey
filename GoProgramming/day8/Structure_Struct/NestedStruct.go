package main

import "fmt"

func main() {
	result := HR{
		details: Author{"Santosh", "TU", 2006},
	}

	// Print the value
	fmt.Println("\nDetails of Author")
	fmt.Println(result)
	fmt.Println("-----------------------------")

	teacherStudent := Teacher{
		name:    "Santosh Shah",
		subject: "Computer Science",
		exp:     2,
		details: Student{"Maneesh Shah", "Computer Science", 2},
	}

	fmt.Println("\nDetails of Teacher")
	fmt.Println("Teacher's name: ", teacherStudent.name)
	fmt.Println("Teacher's subject: ", teacherStudent.subject)
	fmt.Println("Teacher's exprience: ", teacherStudent.exp)

	fmt.Println("\nDetails of Student")
	fmt.Println("Student's name: ", teacherStudent.details.name)
	fmt.Println("Student's branch: ", teacherStudent.details.branch)
	fmt.Println("Student's year: ", teacherStudent.details.year)

}

// TODO: for student and teacher
type Student struct {
	name, branch string
	year         int
}

// creating nested structure
type Teacher struct {
	name, subject string
	exp           int
	details       Student
}

// TODO: for author and HR
type Author struct {
	name, branch string
	year         int
}

type HR struct {
	details Author
}
