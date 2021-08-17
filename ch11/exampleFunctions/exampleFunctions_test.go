package exampleFunctions

import "fmt"

func ExampleLengthRange() {
	fmt.Println(LengthRange("Mihalis"))
	fmt.Println(LengthRange("Mastering Go, 3rd edition!"))
	// Output:
	// 7
	// 7
}
