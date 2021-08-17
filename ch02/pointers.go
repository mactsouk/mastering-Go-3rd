package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)
	// Pointer to f
	fP := &f
	fmt.Println("Memory address of f:", fP)
	fmt.Println("Value of f:", *fP)
	// The value of f changes
	processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f)
	// The value of f does not change
	x := returnPointer(f)
	fmt.Printf("Value of x: %.2f\n", *x)
	// The value of f does not change
	xx := bothPointers(fP)
	fmt.Printf("Value of xx: %.2f\n", *xx)

	// Check for empty structure
	var k *aStructure
	// This is nil because currently k points to nowhere
	fmt.Println(k)
	// Therefore you are allowed to do this:
	if k == nil {
		k = new(aStructure)
	}
	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nil!")
	}
}
