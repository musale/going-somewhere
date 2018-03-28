package main

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat text 5 times", func(t *testing.T) {
		rpt := Repeat("a", 5)
		expected := "aaaaa"
		if rpt != expected {
			t.Errorf("expected %s but got %s", expected, rpt)
		}
	})

	t.Run("allow custom times to repeat", func(t *testing.T) {
		rpt := Repeat("a", 7)
		expected := "aaaaaaa"
		if rpt != expected {
			t.Errorf("expected %s but got %s", expected, rpt)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 5)
	}
}
