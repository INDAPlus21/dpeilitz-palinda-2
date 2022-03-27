// Stefan Nilsson 2013-03-13

// This is a testbed to help you understand channels better.

/*
What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function?
If the statements are switched, the function will terminate before the waitqueue has finished and the program terminates as the goroutines don't have any channels to send to

What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?
The channel will be closed at the end of produce. An error will occurr as the first produce routine will close the channel and the other goroutines won't have a channel to send to

What happens if you remove the statement close(ch) completely?
The channels will still close when the main function ends, ergo the program will still work

What happens if you increase the number of consumers from 2 to 4?
The program will have more threads and therefore be quicker

Can you be sure that all strings are printed before the program stops?
No, as there is no waitgroup that waits until all consumers are finished. Therefore, it is possible that the program terminates before the consume routines finish printing

*/
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 8
	const consumers = 8

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgc := new(sync.WaitGroup)
	wgp.Add(producers)
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}

	wgc.Add(consumers)
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch, wgc)
	}
	wgp.Wait() // Wait for all producers to finish.
	close(ch)
	wgc.Wait() // Wait for all consumers to finish
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string, wg *sync.WaitGroup) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
	wg.Done()
}

// RandomSleep waits for x ms, where x is a random number, 0 < x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}
