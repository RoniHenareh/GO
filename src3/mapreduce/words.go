package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	DataFile = "loremipsum.txt"
	Rutiner  = 8 // ska vara valbart
)

func add(words []string, res chan<- map[string]int) {

	freq := make(map[string]int)

	for _, word := range words {

		freq[word] += 1
	}
	res <- freq
}

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string) map[string]int {

	text = strings.Replace(text, ".", " ", -1)
	text = strings.Replace(text, ",", " ", -1)
	text = strings.ToLower((text))

	words := strings.Fields(string(text)) // splitta

	// now do this cocurrently
	n := len(words)
	ch := make(chan map[string]int)

	lower := 0

	// gör antal go rutiner valbara
	for i := 0; i < Rutiner; i++ {

		upper := n * (i + 1) / Rutiner
		go add(words[lower:upper], ch) // fixa här
		lower = upper

	}

	freqs := make(map[string]int)

	// får info från kanalen

	for i := 0; i < Rutiner; i++ {

		freq1 := <-ch

		for j := range freq1 {
			freqs[j] += freq1[j]

		}

	}

	return freqs

}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data

	// reading code from https://gosamples.dev/read-file/
	data, err := os.ReadFile(DataFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", WordCount(string(data)))

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
