package main

import "fmt"

func main() {
	sliceElem := []int{6, 3, 8, 50, 1, 0}
	target := 5
	output := isElementFound(sliceElem, target)
	if output {
		fmt.Println("Yes ", target, " found!")
	} else {
		fmt.Println("No ", target, " not found!")
	}
}

func isElementFound(sliceElem []int, target int) bool {
	for _, value := range sliceElem {
		if value == target {
			return true
		}
	}

	return false
}
