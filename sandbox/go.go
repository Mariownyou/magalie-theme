package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Hello, World!")
	regexp := regexp.MustCompile(`\d+`)
	symbol := string('%') + string(0x201231ba) + "s"

	fmt.Println(symbol)
	fmt.Println(regexp.MatchString("123"))

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
