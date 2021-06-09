package main

import (
	"fmt"
	"log"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

// Square is a square_num_solution
type Square struct {
	Center Point
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if x <= 0 {
		return nil, fmt.Errorf("Error: x must be a positive number (x was %d)\n", x)
	}

	if y <= 0 {
		return nil, fmt.Errorf("Error: y must be a positive number (y was %d)\n", y)
	}

	if length <= 0 {
		return nil, fmt.Errorf("Error: lenght must be a number greater than 0 (lenght was %d)", length)
	}

	s := &Square{
		Center: Point{x, y},
		Length: length,
	}

	return s, nil
}

func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {

	square, err := NewSquare(1, 1, 4)
	if err != nil {
		log.Fatalf("ERROR: can't create square_num_solution")
	}
	fmt.Printf("square_num_solution init position x=%d & y=%d\n", square.Center.X, square.Center.Y)
	square.Move(1, 2)
	fmt.Printf("square_num_solution moved to position x=%d & y=%d\n", square.Center.X, square.Center.Y)
	fmt.Println("Area of square_num_solution -> ", square.Area())

	square2, err2 := NewSquare(1, 1, 0)
	if err2 != nil {
		log.Fatalf("ERROR: can't create square_num_solution")
	}
	fmt.Println(square2)
}
