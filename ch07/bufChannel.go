package main

import (
	"fmt"
)

func main() {
	// numbers cannot store more than 5 integers
	numbers := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		// This is where the processing takes place
		case numbers <- i * i:
			fmt.Println("About to process", i)
		default:
			fmt.Print("No space for ", i, " ")
		}
	}
	fmt.Println()

	for {
		select {
		case num := <-numbers:
			fmt.Print("*", num, " ")
		default:
			fmt.Println("Nothing left to read!")
			return
		}
	}
}
