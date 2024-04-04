package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for out Application:")

	// comma ok || err err
	input, _ := reader.ReadString('\n')
	//if err != nil {
	//	fmt.Println("An error occurred:", err)
	//	return
	//}
	fmt.Println("Thanks for rating, ", input)
	fmt.Println("Type of this rating is ", input)
}
