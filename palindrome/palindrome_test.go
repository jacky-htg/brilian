package main

import "testing"

func BenchmarkIsPalindromeCompare(b *testing.B) {
	word := "AmanaplanacanalPanama"
	for i := 0; i < b.N; i++ {
		isPalindromeCompare(word)
	}
}

func BenchmarkIsPalindromeReverse(b *testing.B) {
	word := "AmanaplanacanalPanama"
	for i := 0; i < b.N; i++ {
		isPalindromeReverse(word)
	}
}

// Fungsi pengujian dummy
func TestDummy(t *testing.T) {
	// Kosong, hanya untuk memastikan file ini dikenali sebagai file pengujian
}
