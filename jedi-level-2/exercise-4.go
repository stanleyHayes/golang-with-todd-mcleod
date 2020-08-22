package main

import "fmt"

func main() {
	integer := 5
	fmt.Printf("decimal: %d\tbinary: %b\thex: %#x\n", integer, integer, integer)
	integer <<= 1
	fmt.Printf("decimal: %d\tbinary: %b\thex: %#x\n", integer, integer, integer)
}
