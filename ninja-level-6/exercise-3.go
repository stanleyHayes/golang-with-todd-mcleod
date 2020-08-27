package main

import "fmt"

func main() {
	defer deferredFunc()
	first()
}

func first() {
	fmt.Println("Should print first")
}

func deferredFunc() {
	fmt.Println("Called first but deferred")
}