package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

type Shape interface {
	Area() float64
}

func main() {

	s := &Square{4}
	fmt.Println("Area square = ", s.Area())

	c := &Circle{3}
	fmt.Printf("Area circle = %0.4f\n", c.Area())

	shapes := []Shape{s, c}
	sa := sumAreas(shapes)
	fmt.Printf("the sum area of all shapes is = %0.4f ", sa)

}
