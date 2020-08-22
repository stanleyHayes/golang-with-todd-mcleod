package main

import "fmt"

var x int
var y string
var z bool
func main() {
	fmt.Printf("%v %v %v\n", x, y, z)
	s := fmt.Sprintf("\nx: %v\ty: %v\tx: %v\n", x, y, z)
	fmt.Print(s)
}
