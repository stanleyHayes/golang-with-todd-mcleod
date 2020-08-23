package main

import "fmt"

func main() {
	friendFavorites := map[string][]string{
		"afiba": []string{"Java", "C++", "Android Development"},
		"afia":  []string{"DSP", "C++", "Differential Equations"},
		"zeus":  []string{"C", "Software Engineering", "C++"},
	}
	fmt.Println(friendFavorites)
	fmt.Println()
	fmt.Println("-------------------------------------------------")

	for key, friend := range friendFavorites {
		fmt.Printf("This is the record for %v\n", key)
		for index, favorite := range friend {
			fmt.Printf("\t\t---%v\t%v\n", index, favorite)
		}
	}

	fmt.Println("-------------------------------------------------")
	fmt.Println("After adding a record")
	fmt.Println("-------------------------------------------------")

	friendFavorites["akosua"] = []string{"Compiler Construction", "Web Design", "Computer Architecture"}

	for key, friend := range friendFavorites {
		fmt.Printf("This is the record for %v\n", key)
		for index, favorite := range friend {
			fmt.Printf("\t\t---%v\t%v\n", index, favorite)
		}
	}

	fmt.Println("-------------------------------------------------")
	fmt.Println("After deleting a record")
	fmt.Println("-------------------------------------------------")

	delete(friendFavorites, "akosua")

	for key, friend := range friendFavorites {
		fmt.Printf("This is the record for %v\n", key)
		for index, favorite := range friend {
			fmt.Printf("\t\t---%v\t%v\n", index, favorite)
		}
	}

}
