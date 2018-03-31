package main

import "testing"

func TestPerimeter(t *testing.T) {
	perimeter := Perimeter(30.5, 40.8)
	expected := 142.6
	if perimeter != expected {
		t.Errorf("expected %.2f but found %.2f", expected, perimeter)
	}
}

func TestArea(t *testing.T) {
	area := Area(12.0, 3.0)
	expected := 36.0
	if area != expected {
		t.Errorf("expected %.2f but found %.2f", expected, area)
	}
}
