package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string
	Age   int
}

func main() {
	u1 := User{
		First: "Emmanuella",
		Age:   23,
	}

	u2 := User{
		First: "Philomina",
		Age:   22,
	}

	people := []User{u1, u2}
	p, err := json.Marshal(&people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(p))
}
