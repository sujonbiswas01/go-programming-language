package main

import (
	"encoding/json"
	"fmt"
)

type Persons struct {
	Name  string `json:"myName"`
	Email string `json:"email"`
}

func CustomJson() {

	per := Persons{
		Name:  "sujon biswas",
		Email: "sujonbiswas@gmail.com",
	}

	data, err := json.Marshal(per)
	if err != nil {
		fmt.Println("some went wrond", err)
	}
	fmt.Println(data)

}
