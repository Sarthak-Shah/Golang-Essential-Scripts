package main

import (
	"context"
	"fmt"
	"time"
)

//The context package provides a way to manage the lifecycle of Go routines and handle cancellation
//or timeouts.

func worker(ctx context.Context) {
	for {
		select {
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		case <-ctx.Done():
			fmt.Println("Cancelled!")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go worker(ctx)

	time.Sleep(time.Second * 5)
}

/*In this example, the worker Go routine is controlled by a context that provides a timeout of 3 seconds.
After the timeout, the context's Done channel is closed, which triggers the worker to cancel and stop.

These examples provide a basic understanding of Go routines, channels, select, and the context package.
They're crucial tools in building concurrent and parallel programs in Go while ensuring safety and control.*/
