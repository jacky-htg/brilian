package main

import "fmt"

func main() {
	res1 := merge([][]int{
		{1, 2, 5, 4, 3},
		{1, 2, 7, 8},
	})
	fmt.Println(res1)

	res2 := merge([][]int{
		{1, 2, 5, 4, 5},
		{7, 9, 10, 5},
	})
	fmt.Println(res2)

	res3 := merge([][]int{
		{1, 4, 5},
		{7},
	})
	fmt.Println(res3)
}

func merge(data [][]int) []int {
	var merged_map = map[int]int{}
	var merged []int

	for i := 0; i < 2; i++ {
		for j := 0; j < len(data[i]); j++ {
			if merged_map[data[i][j]] == 0 {
				merged_map[data[i][j]] = data[i][j]
				merged = append(merged, data[i][j])
			}
		}
	}

	return merged
}
