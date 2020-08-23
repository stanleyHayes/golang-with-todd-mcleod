package main

import "fmt"

func main() {
	emmanuella := []string{"Emmanuella", "Sangmuah", "Computer Science"}
	afia := []string{"Philomina", "Sarpong", "Biomedical Engieering"}

	friends := [][]string{emmanuella, afia}
	fmt.Println(friends)

	for index, friend := range friends {
		fmt.Printf("Record %v\n", index)
		for i, field := range friend{
			fmt.Printf("%v\t%v\n", i, field)
		}
	}
}
