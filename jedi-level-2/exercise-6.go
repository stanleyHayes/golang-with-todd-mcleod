package main

import "fmt"

const (
	year2017 = iota + 2017
	year2018 = iota + year2017
	year2019 = iota + year2017
	year2020 = iota + year2017
)

func main() {

	fmt.Printf("%v\t%v\t%v\t%v", year2017, year2018, year2019, year2020)
}
