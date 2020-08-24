package main

import "fmt"

func main() {
	value := 666
	fmt.Printf("Address: %v\nValue: %v\n", &value, value)
}
