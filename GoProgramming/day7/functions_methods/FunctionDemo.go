package main

import "fmt"

func main() {

	//TODO: finding area
	//value := area(45, 10)
	//fmt.Println("Area of circle: ", value)

	//value := area(50, 15)
	//fmt.Println(value)

	//fmt.Println(area(50, 20))

	//TODO: swapping values
	//value1 := 10
	//value2 := 20
	//fmt.Printf("Before Swapping value1: %d and value2: %d\n", value1, value2)
	//fmt.Println("----After swapping----")
	//swappedValue(value1, value2)

	//TODO: lets understand call by reference
	value1 := 40
	value2 := 50
	fmt.Printf("Value1: %d and Value2: %d\n", value1, value2)

	fmt.Println("After swapping------------")
	// call by reference
	refSwap(&value1, &value2)
	fmt.Printf("Value1: %d and Value2: %d\n", value1, value2)

}

func refSwap(value1, value2 *int) {
	temp := *value1
	*value1 = *value2
	*value2 = temp
}

func swappedValue(value1, value2 int) {
	temp := value1
	value1 = value2
	value2 = temp
	fmt.Printf("Before Swapping value1: %d and value2: %d\n", value1, value2)
}

func area(length, width int) string {
	value := length * width
	return fmt.Sprintf("Area of circle is: %d", value)
}
