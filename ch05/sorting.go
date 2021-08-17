package main

import (
	"fmt"
	"sort"
)

type Grades struct {
	Name    string
	Surname string
	Grade   int
}

func main() {
	data := []Grades{{"J.", "Lewis", 10}, {"M.", "Tsoukalos", 7},
		{"D.", "Tsoukalos", 8}, {"J.", "Lewis", 9}}

	isSorted := sort.SliceIsSorted(data, func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	})

	if isSorted {
		fmt.Println("It is sorted!")
	} else {
		fmt.Println("It is NOT sorted!")
	}

	sort.Slice(data,
		func(i, j int) bool { return data[i].Grade < data[j].Grade })
	fmt.Println("By Grade:", data)
}
