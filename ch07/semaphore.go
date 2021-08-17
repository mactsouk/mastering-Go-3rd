package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/sync/semaphore"
)

// Maximum number of goroutines
var Workers = 4
var sem = semaphore.NewWeighted(int64(Workers))

func worker(n int) int {
	square := n * n
	time.Sleep(time.Second)
	return square
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need #jobs!")
		return
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Where to store the results
	var results = make([]int, nJobs)

	// Needed by Acquire()
	ctx := context.TODO()

	for i := range results {
		err = sem.Acquire(ctx, 1)
		if err != nil {
			fmt.Println("Cannot acquire semaphore:", err)
			break
		}

		go func(i int) {
			defer sem.Release(1)
			temp := worker(i)
			// No race conditions here - each goroutine writes
			// to a different slice element
			results[i] = temp
		}(i)
	}

	// Acquire all of the tokens
	// This is similar to Wait()
	// It blocks until all workers have finished
	err = sem.Acquire(ctx, int64(Workers))
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range results {
		fmt.Println(k, "->", v)
	}
}
