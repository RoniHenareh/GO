// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {

	questions := make(chan string)
	answers := make(chan string) // kanal för svar

	// TODO: Answer questions.
	// TODO: Make prophecies.
	// TODO: Print answers.

	go func() {
		for {
			question := <-questions
			go answersQuestion(question, answers)
		}
	}()

	go makeProphecies(answers, 30*time.Second)
	go printAnswers(answers)

	return questions
}

// tar en fråga och skickar ett svar till kanalen answers
func answersQuestion(question string, answers chan<- string) {

	svordomar := map[string]string{"fan": "kysser du din mamma med den munnen?", "helvete": "kysser du din mamma med den munnen?", "hora": "kysser du din mamma med den munnen?", "fuck": "kysser du din mamma med den munnen?"}
	frågor := map[string]string{"Kan": "Ska jag behöva göra allt?", "kan": "Ska jag behöva göra allt?", "hjälp": "Ska jag behöva göra allt?", "hjälpa": "Ska jag behöva göra allt?"}
	// kan utvidgas

	// fixa

	for _, ord := range strings.Fields(question) {

		if val, found := svordomar[ord]; found {

			if found {
				answers <- val

			}
			break

		}

		if val, found := frågor[ord]; found {

			if found {
				answers <- val

			}
			break

		}

	}

	time.Sleep(30 * time.Second)
	answers <- "Hmmm, låt mig tänka på saken"

}

// tar ett svar från kanalen och anropar funktionen prophecy
func makeProphecies(answers chan<- string, delay time.Duration) {

	for {

		time.Sleep(delay)

		prophecy("", answers)
		//fmt.Print(prompt)

	}

}

// svarar på frågor och prophecies
func printAnswers(answers <-chan string) {

	for message := range answers {
		fmt.Println("The Oracle has something to say: ")

		delord := strings.Split(message, "")
		for _, ord := range delord {

			fmt.Printf(ord)
			time.Sleep(200 * time.Millisecond)

		}
		fmt.Println("")
		fmt.Print(prompt)

	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.

	//time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark and full of terrors",
		"The sun is bright, dont fly to close to it",
		"This is not a charity, you need to pay first",
	}

	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
