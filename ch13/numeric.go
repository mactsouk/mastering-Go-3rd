package main

import (
	"fmt"
)

type Numeric interface {
    type int, int8, int16, int32, int64, float64
}

func Add[T Numeric](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("4 + 3 =", Add(4,3))
	fmt.Println("4.1 + 3.2 =", Add(4.1,3.2))
}
