package main

import (
	"fmt"
)

func returnNumber() interface{} {
	return 12
}

func main() {
	anInt := returnNumber()
	number := anInt.(int)
	number++
	fmt.Println(number)

	// The next statement will fail because there
	// is no type assertion to get the value:
	// anInt++

	// This fails but the failure is under control
	// because of the `ok` bool variable that tells
	// whether the type assertion is successful or not
	value, ok := anInt.(int64)
	if ok {
		fmt.Println("Type assertion successful: ", value)
	} else {
		fmt.Println("Type assertion failed!")
	}

	// This is successful but dangerous
	// because it does not make sure that
	// the type assertion is successful
	// It just happens to be successful
	i := anInt.(int)
	fmt.Println("i:", i)

	// This will PANIC because `anInt` is not `bool`
	_ = anInt.(bool)
}
