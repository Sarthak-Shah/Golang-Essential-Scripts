package main

import (
	"fmt"
	"time"
)

//Go routines are lightweight threads managed by the Go runtime. They allow you to run functions
//concurrently without blocking the main execution flow.

/*In this example, the printNumbers and printLetters functions are executed concurrently using Go
routines. The main function will wait for a while using time.Sleep to allow the Go routines to finish.*/

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
}

func printLetters() {
	for char := 'a'; char <= 'e'; char++ {
		fmt.Printf("%c ", char)
		time.Sleep(time.Millisecond * 400)
	}
}

func main() {
	go printNumbers()
	go printLetters()

	time.Sleep(time.Second * 3)
}
