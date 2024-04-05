package main

import (
	"fmt"
)

func main() {
	//switch day := 1; day {
	//case 1:
	//	fmt.Println("Sunday")
	//case 2:
	//	fmt.Println("Monday")
	//case 3:
	//	fmt.Println("Tuesday")
	//case 4:
	//	fmt.Println("Wednesday")
	//case 5:
	//	fmt.Println("Thursday")
	//case 6:
	//	fmt.Println("Friday")
	//case 7:
	//	fmt.Println("Saturday")
	//default:
	//	fmt.Println("Invalid Days")
	//}

	// using integer value
	//value := 2
	//switch {
	//case value == 2:
	//	fmt.Println("Two Hello")
	//case value == 3:
	//	fmt.Println("Three Hello")
	//default:
	//	fmt.Println("Invalid Number")
	//}

	// using string value
	//value := "Hi"
	//switch {
	//case value == "Hello":
	//	fmt.Println("This is Hello")
	//case value == "Hi":
	//	fmt.Println("This is Hi")
	//default:
	//	fmt.Println("Invalid Number")
	//}

	// concept ot type switch
	//var value interface{}
	//switch q := value.(type) {
	//case bool:
	//	fmt.Print("Value is of boolean type")
	//case float64:
	//	fmt.Println("value is of float64 type")
	//case int:
	//	fmt.Println("value is of int type")
	//default:
	//	fmt.Printf("value is of type: %T", q)
	//}

	//value := "hello"
	//value := 42
	value := 3.15
	//value := true
	//value := []int {1, 2, 3}
	printType(value)
}

func printType(value interface{}) {
	switch q := value.(type) {
	case bool:
		fmt.Print("Value is of boolean type")
	case float64:
		fmt.Println("value is of float64 type")
	case int:
		fmt.Println("value is of int type")
	default:
		fmt.Printf("value is of type: %T", q)
	}
}
