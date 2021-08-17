package main

import (
	"fmt"
	"time"
)

func printme(x int) {
	fmt.Println("*", x)
	return
}

func main() {
	go func(x int) {
		fmt.Printf("%d ", x)
	}(10)

	go printme(15)

	time.Sleep(time.Second)
	fmt.Println("Exiting...")
}
