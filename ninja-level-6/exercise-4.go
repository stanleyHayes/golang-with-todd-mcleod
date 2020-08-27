package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

func (p person) speak() {
	fmt.Printf("I am %v %v and I am %v years old", p.First, p.Last, p.Age)
}
func main() {

	afiba := person{
		First: "Emmanuella",
		Last: "Sangmuah",
		Age: 23,
	}
	afiba.speak()
}
