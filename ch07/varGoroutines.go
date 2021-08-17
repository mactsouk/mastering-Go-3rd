package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	count := 10
	arguments := os.Args
	if len(arguments) == 2 {
		t, err := strconv.Atoi(arguments[1])
		if err == nil {
			count = t
		}
	}

	fmt.Printf("Going to create %d goroutines.\n", count)

	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
