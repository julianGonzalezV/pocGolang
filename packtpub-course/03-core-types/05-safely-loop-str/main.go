package main

import "fmt"

func main() {

	logLevel := "デバッグ"
	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}

	/// uN BUG que puede pasar es con el mal uso de len
	username := "Sir_King_Über"
	// Length of a string get a wrong answer
	fmt.Println("Bytes:", len(username))

	// Length of a Rune IS THE CORRECT WAY
	fmt.Println("Runes:", len([]rune(username)))
	// Limit to 10 characters
	fmt.Println(string(username[:10]))
	fmt.Println(string([]rune(username)[:10]))
}
