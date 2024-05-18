package main

import "fmt"

func main() {
	var functionScope int = 200 // function-level scope
	fmt.Println("Package Scope:", packageScope)
	fmt.Println("Function Scope:", functionScope)

	if true {
		var blockScope int = 500
		fmt.Println("Block Scope:", blockScope)
	}

	//TODO: Note: Package Level Scope -> Function Level Scope -> Block Level Scope
	//TODO:       "PLS" can be call inside "FLS" can be call inside "BLS"

}

var packageScope = 20
