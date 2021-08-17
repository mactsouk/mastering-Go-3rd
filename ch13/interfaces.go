package main

import (
	"fmt"
)

type Numeric interface {
    type int, int8, int16, int32, int64, float64
}

func Print(s interface{}) {
	// type switch
	switch s.(type) {
	case int:
		fmt.Println(s.(int)+1)
	case float64:
		fmt.Println(s.(float64)+1)
	default:
		fmt.Println("Unknown data type!")
	}
}

func PrintGenerics[T any](s T) {
	fmt.Println(s)
}

func PrintNumeric[T Numeric](s T) {
	fmt.Println(s+1)
}

func main() {
	Print(12)
	Print(-1.23)
	Print("Hi!")
	
	PrintGenerics(1)
	PrintGenerics("a")
	PrintGenerics(-2.33)
	
	PrintNumeric(1)
	PrintNumeric(-2.33)
}
