package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) string {

	for { // oändlighets loop + slepp

		time.Sleep(delay)

		fmt.Println(time.Now().Format("15:04:05"), text)

	}
}

func main() {

	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	go Remind("Time to sleep", 60*time.Second)
	select {}

}
