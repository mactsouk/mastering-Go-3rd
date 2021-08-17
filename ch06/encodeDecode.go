package main

import (
	"encoding/json"
	"fmt"
)

type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"created"`
}

func main() {
	useall := UseAll{Name: "Mike", Surname: "Tsoukalos", Year: 2021}

	// Regular Structure
	// Encoding JSON data -> Convert Go Structure to JSON record with fields
	t, err := json.Marshal(&useall)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %s\n", t)
	}

	// Decoding JSON data given as a string
	str := `{"username": "M.", "surname": "Ts", "created":2020}`
	// Convert string into a byte slice
	jsonRecord := []byte(str)
	// Create a structure variable to store the result
	temp := UseAll{}
	err = json.Unmarshal(jsonRecord, &temp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data type: %T with value %v\n", temp, temp)
	}
}
