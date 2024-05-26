package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Barkkkk!")
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meu Meu!!")
}

type Cow struct{}

func (c Cow) Speak() {
	fmt.Println("Wooh Wooh!")
}

func main() {
	dog := Dog{}
	cat := Cat{}
	cow := Cow{}

	animals := []Animal{dog, cat, cow}
	for _, animal := range animals {
		animal.Speak()
	}
}
