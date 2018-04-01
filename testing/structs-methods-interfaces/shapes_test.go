package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{30.5, 40.8}
	perimeter := Perimeter(rectangle)
	expected := 142.6
	if perimeter != expected {
		t.Errorf("expected %.2f but found %.2f", expected, perimeter)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles area", func(t *testing.T) {
		rectangle := Rectangle{12.0, 3.0}
		area := Area(rectangle)
		expected := 36.0
		if area != expected {
			t.Errorf("expected %.2f but found %.2f", expected, area)
		}
	})

	t.Run("circles area", func(t *testing.T) {
		circle := Circle{10}
		area := Area(circle)
		expected := 314.16
		if area != expected {
			t.Errorf("expected %.2f but found %.2f", expected, area)
		}
	})
}
