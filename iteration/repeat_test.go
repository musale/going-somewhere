package main

import "testing"

func TestRepeat(t *testing.T) {
	rpt := Repeat("a")
	expected := "aaaaa"
	if rpt != expected {
		t.Errorf("expected %s but got %s", expected, rpt)
	}
}
