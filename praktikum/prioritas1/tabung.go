package main

import (
	"fmt"
)

func main() {
	const phi float64 = 3.14

	var r, t int

	fmt.Print("Masukkan jari-jari: ")
	fmt.Scan(&r)

	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&t)

	vol := phi * float64(r*r*t)

	fmt.Printf("%.2f\n", vol)
}
