package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var input string
	fmt.Println("Enter a string:")
	// `%[^\n]` tells Scanf to read everything until a newline is encountered
	fmt.Scan(&input)

	// Count the number of runes in the input string
	numRunes := utf8.RuneCountInString(input)
	costPerChar := 23 // cost per character in cents
	totalCost := numRunes * costPerChar
	dollars := totalCost / 100
	cents := totalCost % 100

	if dollars > 0 {
		fmt.Printf("%d $ %d cents\n", dollars, cents)
	} else {
		fmt.Printf("%d cents\n", cents)
	}
}
