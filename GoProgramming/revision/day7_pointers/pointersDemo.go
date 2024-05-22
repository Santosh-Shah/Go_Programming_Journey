package main

import "fmt"

func addition(num *int) {
	*num = *num + 10
}

func addition1(num int) {
	num = num + 10
}

func getPointer() *int {
	a := 10
	return &a
}

func modifyArray(arr *[5]int) {
	arr[0] = 45
	arr[2] = 13
}

type Person struct {
	name string
	age  int
}

func main() {

	//TODO: basics pointer
	//var a int = 10
	//var ptr *int
	//
	//ptr = &a
	//fmt.Println("Value of a:", a)
	//fmt.Println("Address of a:", &a)
	//fmt.Println("Value of ptr:", ptr)
	//fmt.Println("Value pointed by ptr:", *ptr)

	//TODO: Double Pointer
	//var a int = 20
	//var ptr *int
	//var pptr **int
	//
	//ptr = &a
	//pptr = &ptr
	//
	//fmt.Println("Address of a:", &a)
	//fmt.Println("ptr has:", ptr)
	//
	//fmt.Println("Address of ptr:", &ptr)
	//fmt.Println("pptr has:", pptr)
	//
	//fmt.Println(a, *ptr, **pptr)

	//TODO; Pointers to a function

	//a := 45
	//println("Before:", a)
	//
	//addition1(a)
	//fmt.Println("without Using Pointer function After:", a)
	//
	//addition(&a)
	//fmt.Println("Using Pointer function After:", a)

	//TODO: Returning Pointer from a function
	//ptr := getPointer()
	//fmt.Println("Value:", *ptr)
	//fmt.Println("Address:", ptr)

	//TODO: Pointer to an array as function argument
	//arr := [5]int{1, 2, 3, 4, 5}
	//
	//fmt.Println("Before Array:", arr)
	//modifyArray(&arr)
	//fmt.Println("After Array:", arr)

	//TODO: Pointer to Struct
	//p := Person{"Raju", 20}
	//var ptr = &p
	//
	//fmt.Println("Name:", ptr.name)
	//fmt.Println("Age:", ptr.age)
	//
	//ptr.name = "Rahul"
	//ptr.age = 26
	//
	//fmt.Println("Updated Name:", p.name)
	//fmt.Println("Updated age:", p.age)

	//TODO: Comparing Pointers
	//var a int = 50
	//var b int = 23
	//
	//var ptr1 = &a
	//var ptr2 = &b
	//var ptr3 = &a
	//
	//fmt.Println(ptr1 == ptr2)
	//fmt.Println(ptr1 == ptr3)

	//TODO: Finding the capacity of the pointer
	arr := [5]int{1, 2, 3, 4, 5}
	ptr := &arr
	fmt.Println("Capacity of array:", cap(arr))
	fmt.Println("Capacity of pointer:", cap(*ptr))

	//TODO: Finding the length of pointer
	fmt.Println("Length of array:", len(arr))
	fmt.Println("Length of pointer:", len(*ptr))
}
