package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := `This sentence has five words. Here are five more words. Five-word sentences are fine. But several together become monotonous.`
	lowerText := strings.ToLower(text)

	cleanText := strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return ' '
		}
		return r
	}, lowerText)

	words := strings.Fields(cleanText)
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}

	mostFrequentWord := ""
	maxCount := 0
	for word, count := range wordCounts {
		if count > maxCount {
			mostFrequentWord = word
			maxCount = count
		}
	}

	fmt.Printf("%s", mostFrequentWord)
}
