package main

import "fmt"

func main() {
	numb := foo()
	n, t := bar()
	fmt.Println(n, t, numb)
}

func foo() int {
	return 2
}

func bar() (int, string){
	return 8, "even"
}