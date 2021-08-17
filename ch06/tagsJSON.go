package main

import (
	"encoding/json"
	"fmt"
)

// Ignoring empty fields in JSON
type NoEmpty struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"creationyear,omitempty"`
}

// Removing private fields and ignoring empty fields
type Password struct {
	Name    string `json:"username"`
	Surname string `json:"surname,omitempty"`
	Year    int    `json:"creationyear,omitempty"`
	Pass    string `json:"-"`
}

func main() {
	noempty := NoEmpty{Name: "Mihalis"}
	password := Password{Name: "Mihalis", Pass: "myPassword"}

	// Ignoring empty fields in JSON
	noEmptyVar, err := json.Marshal(&noempty)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("noEmptyVar decoded with value %s\n", noEmptyVar)
	}

	// Removing private fields
	passwordVar, err := json.Marshal(&password)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("password decoded with value %s\n", passwordVar)
	}
}
