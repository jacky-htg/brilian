package main

import "fmt"

func main() {
	var b, t, s, score int

	fmt.Print("Masukkan budget: ")
	fmt.Scan(&b)

	fmt.Print("Masukkan waktu: ")
	fmt.Scan(&t)

	fmt.Print("Masukkan tingkat kesulitan: ")
	fmt.Scan(&s)

	// check budget
	if b >= 0 && b <= 50 {
		score += 4
	} else if b > 50 && b <= 80 {
		score += 3
	} else if b > 80 && b <= 100 {
		score += 2
	} else if b > 100 {
		score += 1
	}

	// check time
	if t >= 0 && t <= 20 {
		score += 10
	} else if t > 20 && t <= 30 {
		score += 5
	} else if t > 30 && t <= 50 {
		score += 2
	} else if t > 50 {
		score += 1
	}

	//difficulty
	if s >= 0 && s <= 3 {
		score += 10
	} else if s > 3 && s <= 6 {
		score += 5
	} else if s > 6 && s <= 10 {
		score += 1
	} else if s > 10 {
		score += 0
	}

	//result
	if score >= 17 && score <= 24 {
		println("High")
	} else if score > 9 && score <= 16 {
		println("Medium")
	} else if score > 2 && score <= 9 {
		println("Low")
	} else if score <= 2 {
		println("Impossible")
	}
}
