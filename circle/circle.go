package main

import (
	"fmt"
	"math"
	"os"
)

// Point is a 2D point
type Point struct {
	X int
	Y int
}

// Circle is an object that has a center and diameter
type Circle struct {
	Center   Point
	Diameter int
}

// Move moves a point object by dx and dy
func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

// Area calculates the area of a circle given its diameter
// area is pi*r^2
func (s *Circle) Area() int {
	return int(math.Pi * math.Pow(float64(s.Diameter/2), 2))
}

// NewCircle creates a circle object using x and y to create a center point and a diameter
func NewCircle(x int, y int, diameter int) (*Circle, error) {
	// Handle errors from func parameters
	if diameter <= 0 {
		return nil, fmt.Errorf("diameter must be >0. It was %d", diameter)
	}

	// Create a point for the circle
	point := Point{
		X: x,
		Y: y,
	}

	// Create Circle
	circle := &Circle{
		Center:   point,
		Diameter: diameter,
	}
	return circle, nil
}

func main() {
	circle, err := NewCircle(3, 5, 12)
	if err != nil {
		fmt.Printf("ERROR: Unable to create the circle. %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully created a circle!\n")
	fmt.Printf("area: %d\n", circle.Area())
	fmt.Printf("diameter: %d\n", circle.Diameter)
	fmt.Printf("center: %d\n", circle.Center)

	fmt.Printf("Moving the circle by -1, -1\n")
	circle.Center.Move(-1, -1)
	fmt.Printf("center: %d\n", circle.Center)
}
