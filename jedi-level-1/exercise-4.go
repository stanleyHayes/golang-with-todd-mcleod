package main

import "fmt"

type integer int

var a integer
var b int
func main() {
	fmt.Println(a)
	fmt.Printf("The type of %v is %T\n", a, a)
	a = 42
	b = int(a)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v, type of y: %T", b, b)
}