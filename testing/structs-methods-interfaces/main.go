package main

func main() {
	Perimeter(30.4, 20.1)
}

// Perimeter gives the perimeter of a 4 sided
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

// Area gives the area of a 4 sided figure
func Area(width, height float64) float64 {
	return width * height
}
