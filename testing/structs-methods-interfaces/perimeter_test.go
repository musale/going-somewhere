package main

import "testing"

func TestPerimeter(t *testing.T) {
	perimeter := Perimeter(30.5, 40.8)
	expected := 142.6
	if perimeter != expected {
		t.Errorf("expected %.2f but found %.2f", expected, perimeter)
	}
}
