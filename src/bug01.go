package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
//THE PROBLEM: Channels are used to send or receive values from goroutines. As no goroutines are used, the channels wait forever for any input/output and the code doesn't work
func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello world!"
	}()

	fmt.Println(<-ch)
}
