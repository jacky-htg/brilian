package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if nilai >= 85 && nilai <= 100 {
		println("A")
	} else if nilai >= 70 && nilai < 85 {
		println("B")
	} else if nilai >= 55 && nilai < 70 {
		println("C")
	} else if nilai >= 40 && nilai < 55 {
		println("D")
	} else if nilai >= 0 && nilai < 40 {
		println("E")
	} else {
		println("Nilai invalid")
	}
}
