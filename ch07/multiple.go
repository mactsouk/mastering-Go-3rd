package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	count := 10
	if len(os.Args) == 1 {
		fmt.Println("Using default value.")
	} else {
		temp, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Using default value.")
		} else {
			count = temp
		}
	}

	fmt.Printf("Going to create %d goroutines.\n", count)
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}
