package main

import "math"

func main() {
	rectangle := Rectangle{30.4, 20.1}
	Perimeter(rectangle)
}

// Shape describes methods for shapes
type Shape interface {
	Area() float64
}

// Rectangle describes a 4 sided figure with 2 opp sides equal
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns area of a rectangle
func (rec Rectangle) Area() float64 {
	return rec.Width * rec.Height
}

// Circle describes a round figure
type Circle struct {
	Radius float64
}

// Area returns area of a circle
func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

// Triangle describes 3 sided figure
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns area of a triangle
func (tri Triangle) Area() float64 {
	return 0.5 * tri.Base * tri.Height
}

// Perimeter gives the perimeter of a 4 sided
func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}
