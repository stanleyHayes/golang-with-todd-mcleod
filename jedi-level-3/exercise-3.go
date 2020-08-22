package main

import "fmt"

func main() {
	initialAlphabet := 65
	for i := 1; i < 27; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", initialAlphabet + i - 1)
		}
	}
}
