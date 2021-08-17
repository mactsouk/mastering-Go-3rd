package main

import "fmt"

func main() {
	aString := "Hello World! €"
	fmt.Println("First character", string(aString[0]))

	// Runes
	// A rune
	r := '€'
	fmt.Println("As an int32 value:", r)
	// Convert Runes to text
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	// Print an existing string as runes
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()

	// String to rune Array
	// myRune := []rune(aString)
	// fmt.Printf("myRune %U\n", myRune)

	// Rune array to string
	// runeArray := []rune{'1', '2', '3'}
	// s := string(runeArray)
	// fmt.Println(s)

	// Print an existing string as characters
	for _, v := range aString {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}
