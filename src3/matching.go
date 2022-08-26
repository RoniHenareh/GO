// http://www.nada.kth.se/~snilsson/concurrency/

package main

import (
	"fmt"
	"sync"
)

// This programs demonstrates how a channel can be used for sending and
// receiving by any number of goroutines. It also shows how  the select
// statement can be used to choose one out of several communications.
func main() {

	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1) // bufferd channel, deadlock wihout it

	wg := new(sync.WaitGroup) // används för att synka alla kanaler
	wg.Add(len(people))       // antal gorutiner att synka

	for _, name := range people {
		go Seek(name, match, wg) // för alla personer anropar vi funtionen Seek
		// där seek skickar eller tar emot ett namn på kanalen "match"
		// whitout "go" we get that the same names are matching
	}

	wg.Wait()

	select {

	case name := <-match:
		fmt.Printf("No one received %sâ€™s message.\n", name)
		//default:
		// There was no pending send operation.

		// behövs för jämnt antal, men inte i vårt fall

	}
}

// Seek either sends or receives, whichever possible, a name on the match
// channel and notifies the wait group when done.
func Seek(name string, match chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
		// Wait for someone to receive my message.
	}
	wg.Done()
}
