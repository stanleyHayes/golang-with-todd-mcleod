package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Printf("Always prints")
	case false:
		fmt.Printf("Never prints")
	}
}
