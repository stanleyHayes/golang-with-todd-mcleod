package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3,4, 5, 6, 7, 8, 9, 0}
	fmt.Printf("%v\n", numbers)
	for key, number := range numbers {
		fmt.Printf("%v\t%v\n", key, number)
	}

	fmt.Printf("%T", numbers)
}
