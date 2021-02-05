package main

import "math"

// Shape interface for rectangle and circle area calculation
type Shape interface {
	Area() float64
}

// Rectangle represents width and height of the rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle represents radius of the circle
type Circle struct {
	Radius float64
}

// CalculatePerimeter calculates of the perimeter and returns
func CalculatePerimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area calculates of the area of the rectangle and returns
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area calculates of the area of the circle and returns
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {

}
