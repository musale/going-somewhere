package main

import "testing"

func TestRepeat(t *testing.T) {
	rpt := Repeat("a", 5)
	expected := "aaaaa"
	if rpt != expected {
		t.Errorf("expected %s but got %s", expected, rpt)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 10)
	}
}
