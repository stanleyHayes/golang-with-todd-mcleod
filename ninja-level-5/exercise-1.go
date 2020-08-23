package main

import "fmt"

type person struct {
	first string
	last string
	flavors []string
}
func main() {
	 afiba := person{
	 	first: "Emmanuella",
	 	last: "Sangmuah",
	 	flavors: []string{"oreo", "strawberry", "vanilla"},
	 }

	 afia := person{
	 	first : "Philomina",
	 	last: "Sarpong",
	 	flavors: []string{"chocolate", "apple", "orange"},
	 }

	 fmt.Println(afiba.first, afiba.last)
	 fmt.Println(afia.first, afia.last)

	 for index, flavor := range afia.flavors{
	 	fmt.Printf("%v\t%v\n", index, flavor)
	 }

	for index, flavor := range afiba.flavors{
		fmt.Printf("%v\t%v\n", index, flavor)
	}

	persons := map[string]person{
		"afiba":afiba,
		"afia": afia,
	}

	fmt.Println(persons["afiba"].first, persons["afiba"].last, persons["afiba"].flavors)
	fmt.Println(persons["afia"].first, persons["afia"].last, persons["afia"].flavors)

	for key, person := range persons{
		fmt.Println(key, person.first, person.last)
		for index, flavor := range person.flavors {
			fmt.Printf("\t\t--%v\t%v\n", index, flavor)
		}
	}
}
