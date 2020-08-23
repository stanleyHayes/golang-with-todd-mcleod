package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	afiba := person{
		first: "Emmanuella",
		last: "Sangmuah",
		age: 23,
	}

	secretAgent1 := secretAgent{
		person: afiba,
		ltk: false,
	}
	fmt.Println(secretAgent1)

	fmt.Printf("First: %v\n", secretAgent1.first)
	fmt.Printf("Last: %v\n", secretAgent1.last)
	//for i := 33; i < 123; i++ {
	//	fmt.Printf("%c\t%#U\n", i, i)
	//}
}
