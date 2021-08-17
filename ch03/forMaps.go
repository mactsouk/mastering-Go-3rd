package main

import "fmt"

func main() {
	// range works with maps as well
	aMap := make(map[string]string)
	aMap["123"] = "456"
	aMap["key"] = "A value"
	for key, v := range aMap {
		fmt.Println("key:", key, "value:", v)
	}

	for _, v := range aMap {
		fmt.Print(" # ", v)
	}
	fmt.Println()
}
