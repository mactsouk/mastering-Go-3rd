package main

import "fmt"

const VAT = 24

type Item struct {
	Description string
	Value       float64
}

func Value(price float64) float64 {
	total := price + price*VAT/100
	return total
}

func main() {
	t := Item{Description: "Keyboard", Value: 100}
	t.Value = Value(t.Value)
	fmt.Println(t)

	tP := &Item{}
	*&tP.Description = "Mouse"
	*&tP.Value = 100
	fmt.Println(tP)
}
