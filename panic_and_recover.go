package main

import (
	"fmt"
)

func safeDivide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	if b == 0 {
		panic("Division by zero")
	}

	return a / b
}

func main() {
	result := safeDivide(10, 2)
	fmt.Println("Result:", result)

	result = safeDivide(8, 0)
	fmt.Println("Result:", result) // This line won't be reached
}
