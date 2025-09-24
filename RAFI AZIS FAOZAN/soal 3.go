package main

import "fmt"

func main() {
	var r, L float64
	fmt.Print("Masukkan Jari- jari lingkaran ")
	fmt.Scan(&r)
	pi := 3.14
	L = pi * r * r
	fmt.Printf("Luas lingkaran: %.1f\n", L)

}
