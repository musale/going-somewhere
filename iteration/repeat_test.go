package main

import "testing"

func TestRepeat(t *testing.T) {
	
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b")
	}
}
