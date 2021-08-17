package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	time.Sleep(3 * time.Second)
	close(b)
}

func C(a, b chan struct{}) {
	<-a
	fmt.Println("C()!")
	close(b)
}

func D(a chan struct{}) {
	<-a
	fmt.Println("D()!")
	wg.Done()
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})
	w := make(chan struct{})

	wg.Add(1)
	go func() {
		D(w)
	}()

	wg.Add(1)
	go func() {
		D(w)
	}()

	go A(x, y)

	wg.Add(1)
	go func() {
		D(w)
	}()

	go C(z, w)
	go B(y, z)

	wg.Add(1)
	go func() {
		D(w)
	}()

	// This triggers the process
	close(x)
	wg.Wait()
}
