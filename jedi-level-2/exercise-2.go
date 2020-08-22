package main

import "fmt"

func main() {
	number1 := 90
	number2 := 50
	a := number1 == number2
	b := number1 <= number2
	c := number1 >= number2
	d := number1 != number2
	e := number1 < number2
	f := number1 > number2

	fmt.Printf("a: %v\nb: %v\nc: %v\nd: %v\ne: %v\nf: %v\n", a, b, c, d, e, f)
}
