package main

import "fmt"

func main() {
	var str string
	var num int

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			if i%2 == 0 {
				str += "I"
			} else {
				str += "O"
			}
		}
	}

	println(str)
}
