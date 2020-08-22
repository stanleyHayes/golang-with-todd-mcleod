package main

import "fmt"

func main() {
	year := 1993
	currentYear := 2020

	for year <= currentYear {
		fmt.Println(year)
		year++
	}
}
