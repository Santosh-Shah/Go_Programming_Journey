package main

import "fmt"

func main() {

	//TODO: Simple for loop
	//for i := 0; i < 4; i++ {
	//	fmt.Println("Vedas College")
	//}

	//TODO: Infinite loop
	//for {
	//	fmt.Println("Vedas College")
	//}

	//TODO: implementing for loop like while loop
	//i := 5
	//for i >= 1 {
	//	fmt.Println("Vedas College")
	//	i--
	//}

	//TODO: use of range in the loop
	//value := []string{"Ram", "Shyam", "Hari"}
	//for i, j := range value {
	//	fmt.Println(i, " : ", j)
	//}

	//TODO: printing unicode of string
	// String as range in the for loop
	//for i, j := range "XabCd" {
	//	fmt.Printf("The index numbers of %U is %d\n", j, i)
	//}

	//TODO: using map in for loop
	//mmap := map[int]string{
	//	10: "Raju",
	//	20: "Ravi",
	//	30: "Rohit",
	//}
	//
	//for key, value := range mmap {
	//	fmt.Println(key, value)
	//}

	//TODO: using channel in for loop
	chnl := make(chan int)
	go func() {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()

	for i := range chnl {
		fmt.Println(i)
	}
}
