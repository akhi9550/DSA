package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func minimumPushes(word string) int {
	result := 0
	l := len(word)
	for i := 0; i < l; i++ {
		result += ((i / 8) + 1)
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a word: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	if len(word) == 0 {
		fmt.Println("Please enter a valid word.")
		return
	}

	pushes := minimumPushes(word)
	fmt.Printf("The minimum number of pushes for the word '%s' is: %d\n", word, pushes)
}
