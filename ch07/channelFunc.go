package main

import (
	"fmt"
	"time"
)

func printer(ch chan<- bool) {
	ch <- true
}

func writeToChannel(c chan<- int, x int) {
	fmt.Println("1", x)
	c <- x
	fmt.Println("2", x)
}

func f2(out <-chan int, in chan<- int) {
	x := <-out
	fmt.Println("Read (f2):", x)
	in <- x
	return
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
	fmt.Println("Read:", <-c)
	time.Sleep(1 * time.Second)
	close(c)

	c1 := make(chan int, 1)
	c2 := make(chan int, 1)

	// Write to channel
	c1 <- 5
	f2(c1, c2)
	// Read from channel
	fmt.Println("Read (main):", <-c2)
}
