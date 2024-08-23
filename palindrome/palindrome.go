package main

import (
	"strings"
)

// Fungsi untuk membalik string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Pendekatan 1: Membandingkan karakter
func isPalindromeCompare(word string) bool {
	word = strings.ToLower(word)
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}

// Pendekatan 2: Membalik string
func isPalindromeReverse(word string) bool {
	word = strings.ToLower(word)
	reversedWord := reverseString(word)
	return word == reversedWord
}
