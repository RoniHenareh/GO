package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	ord := strings.Fields(s) // split

	dic := make(map[string]int)

	for i := 0; i < len(ord); i++ {
		dic[ord[i]] += 1
	}

	return dic

}

func main() {
	wc.Test(WordCount)
}

// f("A man a plan a canal panama.") =
// map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}
