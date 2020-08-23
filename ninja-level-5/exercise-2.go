package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	truck1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "purple",
		},
		luxury: true,
	}

	fmt.Printf("\n\n--------------------\ndoors count: %v\ncolor: %v\nLuxurious?: %v\n\n",sedan1.doors, sedan1.color, sedan1.luxury)

	fmt.Printf("\n\n--------------------\ndoors count: %v\ncolor: %v\nFour Wheel?: %v",truck1.doors, truck1.color, truck1.fourWheel)
}
