package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "Emmanuella",
		Last: "Sangmuah",
		Age: 23,
		Sayings: []string{"Ebefa", "God no go shame us", "Obi nnim oberempon ahyese"},
	}

	users := []user{u1}

	enc := json.NewEncoder(os.Stdout)
	err := enc.Encode(&users)
	if err != nil {
		fmt.Println(err)
	}
}
