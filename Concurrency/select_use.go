package main

import (
	"fmt"
	"time"
)

//The select statement allows you to wait on multiple channel operations. It's often used to implement
// non-blocking channel communication.

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- 2
	}()

	select {
	case val := <-ch1:
		fmt.Println("Received from ch1:", val)
	case val := <-ch2:
		fmt.Println("Received from ch2:", val)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout!")
	}
}

/*In this example, the select statement waits for events on multiple channels. Whichever channel
receives data first will be selected, or if none receive data within the specified time, a timeout
occurs.*/
