package main

import "fmt"

func removeDuplicates(arr []int) []int {
	unique := make(map[int]bool)
	result := []int{}

	for _, value := range arr {
		if unique[value] != true {
			unique[value] = true
			result = append(result, value)
		}
	}

	return result
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 5}
	uniqueArr := removeDuplicates(arr)
	fmt.Println("Array without duplicates:", uniqueArr)
}
