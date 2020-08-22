package main

import "fmt"

func main() {

	birthYear := 1993
	currentYear := 2020
	for {
		fmt.Println(birthYear)
		if birthYear >=  currentYear {
			break
		}
		birthYear++
	}

}
