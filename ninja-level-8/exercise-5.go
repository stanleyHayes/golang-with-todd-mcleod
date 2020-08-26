package main

import (
	"fmt"
	"sort"
)

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []person

func (ba ByAge) Swap(i int, j int) { ba[i], ba[j] = ba[j], ba[i] }
func (ba ByAge) Len() int          { return len(ba) }
func (ba ByAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }
func main() {
	afibs := person{
		First:   "Emmanuella",
		Last:    "Sangmuah",
		Age:     23,
		Sayings: []string{"Ebefa", "God no go shame us", "Obi nnim oberempon ahyese"},
	}

	afia := person{
		First:   "Philomina",
		Last:    "Sarpong",
		Age:     19,
		Sayings: []string{"I", "Don't", "Even know"},
	}

	sweethearts := []person {afibs, afia}

	for i, p := range sweethearts{
		fmt.Println(i, p)
	}

	sort.Sort(ByAge(sweethearts))

	for i, p := range sweethearts{
		fmt.Println(i, p)
	}
}
