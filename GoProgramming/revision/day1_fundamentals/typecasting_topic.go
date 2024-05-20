package main

import "fmt"

func main() {
	var a int = 42
	var b float64 = float64(a)
	var c int = int(b)
	fmt.Println(a, b, c)
}
