package main

import "fmt"

func main() {
	var str string

	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	//fmt.Println("Len: ", len(str))
	//fmt.Println(str[2:])
	//fmt.Println(str[:3])
	output := reversedString(str)
	fmt.Println("reverseString:", output)
}

func reversedString(str string) string {
	if len(str) == 0 {
		return str
	}

	reversed := reversedString(str[1:]) + string(str[0])
	return reversed
}
