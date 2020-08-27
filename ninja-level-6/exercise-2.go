package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(foo1(numbers...))
	fmt.Println(bar1(numbers))

}

func foo1(numbers ...int) int {
	sum := 0
	for _, number := range numbers{
		sum += number
	}
	return sum
}

func bar1(numbers []int) int {
	sum := 0
	for _, value := range numbers{
		sum += value
	}
	return sum
}