package coverage

import "fmt"

func f1() {
	if true {
		fmt.Println("Hello!")
	} else {
		fmt.Println("Hi!")
	}
}

func f2(n int) int {
	if n >= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return f2(n-1) + f2(n-2)
	}
}
