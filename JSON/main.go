package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("Learning JSON")

	person := Person{Name: "Amith", Age: 34, IsAdult: true}
	fmt.Println("Person data is:", person)

	// Convert to JSON using Marshal
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}
	fmt.Println("JSON data is:", string(jsonData))

	// Unmarshal the JSON data
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}
	fmt.Println("Unmarshalled person data is:", personData)
}
