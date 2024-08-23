package main

import (
	"fmt"
)

func power(x, n int) int {
	result := 1
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n /= 2
	}
	/*for i := 0; i < n; i++ {
		result *= x
	}*/
	return result
}

func main() {
	x := 2
	n := 3
	result := power(x, n)
	fmt.Printf("%d^%d = %d\n", x, n, result)
}
