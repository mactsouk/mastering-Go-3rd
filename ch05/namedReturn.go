package main

import (
	"fmt"
	"os"
	"strconv"
)

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
		return min, max
	}

	min = x
	max = y
	return
}

func main() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("The program needs at least 2 arguments!")
		return
	}

	// No checking here - we trust the user!!
	a1, _ := strconv.Atoi(arguments[1])
	a2, _ := strconv.Atoi(arguments[2])

	fmt.Println(minMax(a1, a2))
	mi, ma := minMax(a1, a2)
	fmt.Println(mi, ma)
}
