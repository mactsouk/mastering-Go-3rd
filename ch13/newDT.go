package main

import (
	"fmt"
	"errors"
)

type TreeLast[T any] []T

func (t TreeLast[T]) replaceLast(element T) (TreeLast[T], error) {
	if len(t) == 0 {
		return t, errors.New("This is empty!")
	}
	
	t[len(t) - 1] = element
	return t, nil
}

func main() {
	tempStr := TreeLast[string]{"aa", "bb"}
	fmt.Println(tempStr)
	tempStr.replaceLast("cc")
	fmt.Println(tempStr)

	tempInt := TreeLast[int]{12, -3}
	fmt.Println(tempInt)
	tempInt.replaceLast(0)
	fmt.Println(tempInt)
}
