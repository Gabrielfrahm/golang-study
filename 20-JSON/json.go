package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	person1 := Person{
		Name:  "Gabriel",
		Email: "gabriel@teste.com",
	}

	fmt.Println(person1)
	personJSON, err := json.Marshal(person1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(personJSON))

	var person2 Person

	if err = json.Unmarshal([]byte(personJSON), &person2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(person2)
}
