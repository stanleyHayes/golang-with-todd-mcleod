package main

import (
	"encoding/json"
	"fmt"
)

type UserStruct struct {
	First string
	Age int
}
func main() {
	people := []UserStruct{}
	result := `[{"First": "Emmanuella", "Age": 23}, {"First": "Philomina", "Age": 22}]`
	err := json.Unmarshal([]byte(result), &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

}
