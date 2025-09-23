package main

import "fmt"

func HitungLuasLingkaran(r float64) float64 {
	const pi float64 = 3.14159
	return pi * (r * r)
}

func main() {
	var r float64

	fmt.Print("Masukkan r: ")
	fmt.Scan(&r)

	result := HitungLuasLingkaran(r)

	fmt.Printf("Luas Lingkaran: %.1f\n", result)
}
