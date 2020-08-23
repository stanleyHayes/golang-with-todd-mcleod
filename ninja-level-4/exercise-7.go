package main

import "fmt"

func main() {
	states := make([]string, 50, 50)
	states = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", ""}
	fmt.Println(len(states), cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Printf("%v\t%v\n", i, states[i])
	}
}
