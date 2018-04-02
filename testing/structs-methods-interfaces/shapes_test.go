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
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, areaTest := range areaTests {
		got := areaTest.shape.Area()
		if got != areaTest.want {
			t.Errorf("expected %.2f but got %.2f", areaTest.want, got)
		}
	}
}
