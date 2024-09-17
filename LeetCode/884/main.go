package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	uncommon := []string{}
	maps := make(map[string]int)
	words1 := strings.Split(s1, " ")
	words2 := strings.Split(s2, " ")
	for _, word := range words1 {
		maps[word]++
	}
	for _, word := range words2 {
		maps[word]++
	}
	for word, count := range maps {
		if count == 1 {
			uncommon = append(uncommon, word)
		}
	}

	return uncommon
}

func main() {
	sentence1 := "the quick brown fox"
	sentence2 := "the lazy dog"
	result := uncommonFromSentences(sentence1, sentence2)
	fmt.Println("Uncommon words:", result)
}
