package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

//In Go, you define methods with a receiver parameter, which is a variable of the type on which
//the method operates. The receiver is placed in front of the method name in the method definition.

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s %s\n", p.FirstName, p.LastName)
}

func main() {
	person := Person{
		FirstName: "Sarthak",
		LastName:  "Shah",
	}

	person.Greet()
}
