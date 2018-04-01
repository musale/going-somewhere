package main

func main() {
	rectangle := Rectangle{30.4, 20.1}
	Perimeter(rectangle)
}

// Rectangle describes a 4 sided figure with 2 opp sides equal
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle describes a round figure
type Circle struct {
	Radius float64
}

// Perimeter gives the perimeter of a 4 sided
func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

// Area gives the area of a 4 sided figure
func Area(rec Rectangle) float64 {
	return rec.Width * rec.Height
}
