package main

import (
	"fmt"
	"runtime/metrics"
	"sync"
	"time"
)

func main() {
	const nGo = "/sched/goroutines:goroutines"

	// A slice for getting metric samples
	getMetric := make([]metrics.Sample, 1)
	getMetric[0].Name = nGo

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(4 * time.Second)
		}()

		// Get actual data
		metrics.Read(getMetric)
		if getMetric[0].Value.Kind() == metrics.KindBad {
			fmt.Printf("metric %q no longer supported\n", nGo)
		}

		mVal := getMetric[0].Value.Uint64()
		fmt.Printf("Number of goroutines: %d\n", mVal)
	}

	wg.Wait()

	metrics.Read(getMetric)
	mVal := getMetric[0].Value.Uint64()
	fmt.Printf("Before exiting: %d\n", mVal)
}
