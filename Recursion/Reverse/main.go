package main

import "fmt"

func Reverse(str string) string {
	if len(str) <= 1 {
		return str
	}
	first := str[0]
	last := str[len(str)-1]
	revese := string(last) + Reverse(str[1:len(str)-1]) + string(first)
	return revese
}

func main() {
	String := "akhil"
	result := Reverse(String)
	fmt.Println("Revesed :-", result)
}
