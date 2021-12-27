package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	var min, max float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}

		// As pointed out by a reader of the book:
		//
		// When the first argument is non-parseable, the `min/max` variables are not initialized
		// to the first parseable argument (but to `0` by default).
		// This would lead to bug when all the parseable arguments share the same sign.
		if i == 1 {
			min = n
			max = n
			continue
		}

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
