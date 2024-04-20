package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Define the data
	names := []string{"brian", "habib", "malik"}
	ages := []int{25, 25, 24}
	hobbies := []string{"hiking", "touring", "traveling"}

	// Create an array to hold person structs
	var people []Person

	// Populate the array with data
	for i := 0; i < len(names); i++ {
		person := Person{
			Name:  names[i],
			Age:   ages[i],
			Hobby: hobbies[i],
		}
		people = append(people, person)
	}

	// Encode the array of structs to JSON
	jsonData, err := json.MarshalIndent(people, "", "    ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Print the JSON
	fmt.Println(string(jsonData))
}

// Define a struct to hold person information
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Hobby string `json:"hobby"`
}
