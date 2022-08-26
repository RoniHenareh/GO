package main

import (
	"fmt"
)

// I want this program to print "Hello world!", but it doesn't work.
func main() {

	ch := make(chan string, 1) // längd ger buffered

	// alt 2
	//ch := make(chan string)
	//go func() {

	//ch <- "hello world!" }()

	//The program will get stuck on the channel send operation
	// waiting forever for someone to read the value.

	//Sends to a buffered channel block only when the buffer is full.
	//Receives block when the buffer is empty.

	ch <- "Hello world!"

	fmt.Println(<-ch)

}
