package main

import "fmt"

type person struct {
	name string
	age int
}

func changeMe(p *person) *person{
	(*p).name = "Stanley"
	(*p).age = 45
	return p
}
func main() {
	afiba := person{
		name: "Emmanuella",
		age: 23,
	}
	ella := changeMe(&afiba)
	fmt.Println(ella.name, ella.age)
}
