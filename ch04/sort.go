package main

import (
	"fmt"
	"sort"
)

type S1 struct {
	F1 int
	F2 string
	F3 int
}

// We want to sort S2 records based on the value of F3.F1
// Which is S1.F1 as F3 is an S1 structure
type S2 struct {
	F1 int
	F2 string
	F3 S1
}

type S2slice []S2

// Implementing sort.Interface for S2slice
func (a S2slice) Len() int {
	return len(a)
}

// What field to use for comparing
func (a S2slice) Less(i, j int) bool {
	return a[i].F3.F1 < a[j].F3.F1
}

func (a S2slice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	data := []S2{
		S2{1, "One", S1{1, "S1_1", 10}},
		S2{2, "Two", S1{2, "S1_1", 20}},
		S2{-1, "Two", S1{-1, "S1_1", -20}},
	}
	fmt.Println("Before:", data)
	sort.Sort(S2slice(data))
	fmt.Println("After:", data)

	// Reverse sorting works automatically
	sort.Sort(sort.Reverse(S2slice(data)))
	fmt.Println("Reverse:", data)
}
