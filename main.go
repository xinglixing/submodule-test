package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello!")
}

func sayGoodbye() {
	fmt.Println("Goodbye!")
}

func sayHi() {
	fmt.Println("Hi!")
}

func main() {

	sayHello()

	fmt.Println("Hello, World!")

	sayHi()

	sayGoodbye()
}
