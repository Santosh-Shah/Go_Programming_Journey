package main

import "fmt"

func main() {

	//TODO This is Anonymous function
	//func() {
	//	fmt.Println("Welcome! to GeeksForGeeks")
	//}()

	//TODO: assigning an anonymous function into variable
	//value := func() {
	//	fmt.Println("assigning an anonymous function into variable")
	//}

	// correct
	//value()

	// correct
	//duplicateValue := value
	//duplicateValue()

	//TODO: passing arguments into anonymous function
	//func(valueStr string) {
	//	fmt.Println(valueStr)
	//}("Vedas College")

	//TODO: passing an anonymous function as an argument into other function.
	value := func(value1, value2 string) string {
		return value1 + value2 + "Geeks"
	}

	collegeName(value)

}

func collegeName(value func(value1, value2 string) string) {
	fmt.Println(value("Geeks", "For"))
}
