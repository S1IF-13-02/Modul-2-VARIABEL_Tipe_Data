package main

import "fmt"

func main() {
	var r float64
	const PI = 3.14159

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&r)

	luas := PI * r * r

	fmt.Printf("Luas lingkaran adalah %.1f\n", luas)
}