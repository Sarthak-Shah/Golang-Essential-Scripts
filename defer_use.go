package main

import "fmt"

/*
In Go, the defer statement is used to schedule a function call to be executed after the surrounding function returns,
regardless of whether the return is due to a normal flow or a panic. This feature can be extremely useful in ensuring
that resources are properly cleaned up, files are closed, connections are closed, etc., even in the presence of panics.

When a function call is deferred, the arguments to the call are evaluated immediately, but the actual function call is
postponed until the surrounding function returns.

The defer statement is often used in combination with panic and recover to ensure that certain cleanup operations are
performed even in the event of a panic.
*/

func main() {
	defer fmt.Println("Deferred call executed")
	//The defer fmt.Println("Deferred call executed") statement is scheduled to be executed after the main function returns.

	fmt.Println("Start")

	// Simulating a panic
	panic("Something went wrong")

	fmt.Println("End") // This line won't be reached
}
