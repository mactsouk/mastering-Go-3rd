package main

import (
	"runtime"
)

func main() {
	var N = 80000000
	split := make([]map[int]int, 2000)
	for i := range split {
		split[i] = make(map[int]int)
	}
	for i := 0; i < N; i++ {
		value := int(i)
		split[i%2000][value] = value
	}
	runtime.GC()
	_ = split[0][0]
}
