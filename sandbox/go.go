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
