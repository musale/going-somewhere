package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{30.5, 40.8}
	perimeter := Perimeter(rectangle)
	expected := 142.6
	if perimeter != expected {
		t.Errorf("expected %.2f but got %.2f", expected, perimeter)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, areaTest := range areaTests {
		t.Run(areaTest.name, func(t *testing.T) {
			got := areaTest.shape.Area()
			if got != areaTest.hasArea {
				t.Errorf("%#v expected %.2f but got %.2f", areaTest.name, areaTest.hasArea, got)
			}
		})

	}
}
