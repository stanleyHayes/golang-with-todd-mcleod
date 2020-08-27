package main

import (
	"fmt"
	"math"
)

type square struct {
	Side float64
}

type circle struct {
	Radius float64
}

func (s square) area() float64 {
	return s.Side * s.Side
}

func (c circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	unitCircle := circle{Radius: 1}
	unitSquare := square{Side: 1}
	fmt.Printf("Area of square with radius %v is %v\n", unitSquare.Side, unitSquare.area())
	fmt.Printf("Area of circle with radius %v is %v", unitCircle.Radius, unitCircle.area())
}
