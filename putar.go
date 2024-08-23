package main

import (
	"fmt"
)

func rotateString(input string) string {
	n := len(input)
	result := make([]byte, n)
	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = input[left]
			left++
		} else {
			result[i] = input[right]
			right--
		}
	}
	return string(result)
}

func main() {
	inputs := []string{"alta", "alterra", "sepulsa", "indonesiarayayangmaju"}
	for _, input := range inputs {
		fmt.Printf("Input: %s, Output: %s\n", input, rotateString(input))
	}
}
