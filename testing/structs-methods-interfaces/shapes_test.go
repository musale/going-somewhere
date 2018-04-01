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
	t.Run("rectangles area", func(t *testing.T) {
		rectangle := Rectangle{12.0, 3.0}
		area := rectangle.Area()
		expected := 72.0
		if area != expected {
			t.Errorf("expected %.2f but got %.2f", expected, area)
		}
	})

	t.Run("circles area", func(t *testing.T) {
		circle := Circle{10}
		area := circle.Area()
		expected := 314.1592653589793
		if area != expected {
			t.Errorf("expected %f but got %f", expected, area)
		}
	})
}
