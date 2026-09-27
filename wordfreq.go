package main

import "fmt"

// WordCount represents a word and how many times it appears.
type WordCount struct {
	Word  string
	Count int
}

// TopN takes a text string and an integer n, and returns the top n most
// frequent words. Results must be sorted by count descending. Words with
// the same count must be sorted alphabetically ascending.
//
// Words are case-insensitive ("Go" and "go" count as the same word).
// Punctuation attached to words should be ignored ("hello," == "hello").
func TopN(text string, n int) []WordCount {
	// TODO: implement this function
	return nil
}

func main() {
	text := "the quick brown fox jumps over the lazy dog the fox"

	results := TopN(text, 3)
	for _, wc := range results {
		fmt.Printf("%s: %d\n", wc.Word, wc.Count)
	}
}
