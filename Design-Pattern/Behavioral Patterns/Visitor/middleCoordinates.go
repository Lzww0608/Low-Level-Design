package main

import "fmt"

type MiddleCoordinates struct {
	x int
	y int
}

func (m *MiddleCoordinates) visitForSquare(s *Square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (m *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (m *MiddleCoordinates) visitForRectangle(r *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}