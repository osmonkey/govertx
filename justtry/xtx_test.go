package main

import (
	"testing"
)

func BenchmarkWp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Wp()
	}
}

func BenchmarkNp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Np()
	}
}
