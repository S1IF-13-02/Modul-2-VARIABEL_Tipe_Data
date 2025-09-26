package main

import "fmt"

func main() {
	var (
		r, L float64
		π    = 3.14
	)
	fmt.Print("Jari-jari = ")
	fmt.Scan(&r)
	L = π * r * r
	fmt.Println("Luas lingkaran =", L)

}
