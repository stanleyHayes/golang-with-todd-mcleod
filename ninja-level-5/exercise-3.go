package main

import "fmt"

type member struct {
	first string
	age int
}
func main() {
	edwin := member{
		first: "Edwin",
		age: 29,
	}

	winston := member{
		first: "Winston",
		age: 28,
	}
 	family := struct {
		name string
		location string
		members []member
	}{
 		name: "Hayford",
 		location: "Cape Coast, Ghana",
 		members: []member{edwin, winston},
	}

	fmt.Println(family)
}
