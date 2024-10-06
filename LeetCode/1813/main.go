package main

import (
    "fmt"
    "strings"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    if sentence1 == sentence2 {
        return true
    }

    maxWords, minWords := MaxMin(sentence1, sentence2)
    
    start, end := 0, 0
    for start < len(minWords) {
        if minWords[start] != maxWords[start] {
            break
        }
        start++
    }

    for end < len(minWords) {
        if minWords[len(minWords)-1-end] != maxWords[len(maxWords)-1-end] {
            break
        }
        end++
    }

    return start + end >= len(minWords)
}

func MaxMin(s1, s2 string) ([]string, []string) {
    s1Words := strings.Split(s1, " ")
    s2Words := strings.Split(s2, " ")
    if len(s1Words) > len(s2Words) {
        return s1Words, s2Words
    }
    return s2Words, s1Words
}

func main() {
    sentence1 := "my name is Akhil"
    sentence2 := "my name is"

    if areSentencesSimilar(sentence1, sentence2) {
        fmt.Println("The sentences are similar.")
    } else {
        fmt.Println("The sentences are not similar.")
    }
}
