package main

import "fmt"

func main() {
	//TODO: break statement
	//for i := 0; i < 10; i++ {
	//	if i == 4 {
	//		break
	//	}
	//	fmt.Println("My name is Khan:", i)
	//}

	//TODO: continue statement
	//for i := 0; i < 10; i++ {
	//	if i%2 == 0 {
	//		continue
	//	}
	//	fmt.Println("Vedas College:", i)
	//}

	//TODO: switch statement
	//day := 7
	//
	//switch day {
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
	//	fmt.Println("Please enter a valid time from 1 - 7 or exact day")
	//}

	//TODO: Switch with Multiple Cases
	//day := "Monday"
	//
	//switch day {
	//case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	//	fmt.Println("It's a week day")
	//case "Saturday", "Sunday":
	//	fmt.Println("It's the weekend")
	//default:
	//	fmt.Println("Not a valid day")
	//}

	//TODO: Switch with condition
	num := 0
	switch {
	case num < 5 && num > 0:
		fmt.Println("num is less than 5")
	case num < 10 && num > 5:
		fmt.Println("num is greater than 5")
	case num > 10:
		fmt.Println("num is greater than 10")
	default:
		fmt.Println("num is equal to 0")
	}
}
