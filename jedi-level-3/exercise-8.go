package main

import "fmt"

func main() {
	var favSport string
	fmt.Printf("Enter your favorite sport: ")
	fmt.Scanf("%s", &favSport)

	switch favSport {
	case "football", "FOOTBALL", "Football":
		fmt.Printf("Welcome to the real game")
	case "basketball", "BASKETBALL", "Basketball":
		fmt.Printf("Well, people do enjoy you but not me")
	default:
		fmt.Printf("Are you even someone's favorit sport")
	}
}
