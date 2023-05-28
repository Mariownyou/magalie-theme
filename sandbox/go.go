package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	for i := 0; i < 10; i++ {
		println(i)
	}
}

func SomeFunc(word string) (string, error) {
	return word, nil
}

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s\n", p.Name)
}

func NewPerson() Person {
	return Person{
		Name: "John",
		Age:  30,
	}
}
