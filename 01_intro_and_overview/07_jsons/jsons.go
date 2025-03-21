package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
	Age        int    `json:"age"`
	LikesDogs  bool   `json:"likes_dogs"`
}

func main() {
	people := `
[
	{
		"name": "Bartek",
		"second_name": "C",
		"age": 27,
		"likes_dogs": true
	},
	{
		"name": "Jonas",
		"second_name": "D",
		"age": 30,
		"likes_dogs": false
	}
]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(people), &unmarshalled)
	if err != nil {
		log.Println("Something went wrong", err)
	}
	log.Printf("Unmarshalled verison: %v, type: %T", unmarshalled, unmarshalled)

	person1 := Person{"Aga", "D", 28, true}
	person2 := Person{"SomeName", "A", 13, false}

	var arrayOfPeople []Person
	arrayOfPeople = append(arrayOfPeople, person1, person2)

	marshalled, err := json.MarshalIndent(arrayOfPeople, "", " .   ")
	if err != nil {
		log.Println("Went wrong", err)
	}
	fmt.Println(string(marshalled))

}
