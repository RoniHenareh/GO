package main

import (
	"fmt"
	"time"
)

// This program should go to 11, but it seemingly only prints 1 to 10.
func main() {

	ch := make(chan int)
	go Print(ch) // oklart

	for i := 1; i <= 11; i++ {
		ch <- i // skickar på kanalen
	}

	//PROBLEM: the for statement can be faster than the the channel reading,
	//which result in closing the channel before all values have been read.
	time.Sleep(time.Second)

	close(ch) // return
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.

func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
}
