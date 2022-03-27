package main

import (
	"fmt"
	"sync"
	"time"
)

// This program should go to 11, but it seemingly only prints 1 to 10.
//THE PROBLEM: The main function finishes before the print function, which closes the program prematurely.
//SOLUTION: Force the main function to wait until Print has finished executing
var wait = new(sync.WaitGroup)

func main() {
	ch := make(chan int)

	wait.Add(1)

	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	wait.Wait()
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
	wait.Done()
}
