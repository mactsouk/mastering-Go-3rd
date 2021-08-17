package main

import "fmt"

type Secret struct {
	SecretValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

func Teststruct(x interface{}) {
	// type switch
	switch T := x.(type) {
	case Secret:
		fmt.Println("Secret type")
	case Entry:
		fmt.Println("Entry type")
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func Learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf("Data type: %T\n", T)
	}
}

func main() {
	A := Entry{100, "F2", Secret{"myPassword"}}
	Teststruct(A)
	Teststruct(A.F3)
	Teststruct("A string")

	Learn(12.23)
	Learn('€')
}
