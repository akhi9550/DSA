package main

import "fmt"

var pickedNumber int

func guess(num int) int {
	if num > pickedNumber {
		return -1
	} else if num < pickedNumber {
		return 1
	}
	return 0
}

func guessNumber(n int) int {
	low := 1
	high := n

	for low <= high {
		myGuess := (high + low) / 2
		result := guess(myGuess)
		if result == 0 {
			return myGuess
		} else if result > 0 {
			low = myGuess + 1
		} else {
			high = myGuess - 1
		}
	}
	return -1
}

func main() {
	pickedNumber = 6
	n := 10
	fmt.Printf("Picked number: %d, Found number: %d\n", pickedNumber, guessNumber(n))
}
