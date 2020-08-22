package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanf("%v", &age)

	fmt.Printf("%v %T", age, age)
	if age < 18 {
		fmt.Printf("You are %v years old and hence cannot vote", age)
	} else if age == 18 {
		fmt.Printf("You almost lost your vote")
	}else {
		fmt.Println("Go ahead and cast your vote")
	}
}
