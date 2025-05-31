package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func main() {
	shapes := []Shape{
		Circle{Radius: 30},
		Triangle{Base: 10, Height: 15},
	}

	for _, shape := range shapes {
		fmt.Printf("Area: %f\n", shape.Area())
	}

}
