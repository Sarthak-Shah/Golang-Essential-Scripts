package main

import (
	"fmt"
	"time"
)

//Channels are used to communicate and synchronize between Go routines.
//They allow safe data exchange.

func sendData(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send data to channel
		time.Sleep(time.Millisecond * 50)
	}
	close(ch) // Close the channel after sending data
}

func main() {
	ch := make(chan int)
	go sendData(ch)

	for num := range ch { // Receive data from channel
		fmt.Println("Received:", num)
	}

	fmt.Println("Channel closed")
}

/*In this example, the sendData function sends integers to the channel ch,
and the main function receives and prints those integers using a for loop. The channel is closed after
all data is sent.*/
