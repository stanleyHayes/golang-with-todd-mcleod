package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Printf("%T\t%v\n", numbers, numbers)

	for index, value := range numbers {
		fmt.Printf("%v\t%v\n", index, value)
	}
}
