package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(primeSum([]int{1, 2, 4, 5, 12, 19, 30}))
	fmt.Println(primeSum([]int{45, 17, 44, 67, 11}))
	fmt.Println(primeSum([]int{15, 12, 9, 10}))
}

func primeSum(numbers []int) int {
	var sum int

	for i := 0; i < len(numbers); i++ {
		flag := 0
		if numbers[i] >= 2 {
			for j := 2; j <= int(math.Sqrt(float64(numbers[i]))); j++ {
				if numbers[i]%j == 0 {
					flag = 1
					break
				}
			}
		} else {
			flag = 1
		}

		if flag == 0 {
			sum += numbers[i]
		}
	}

	return sum
}
