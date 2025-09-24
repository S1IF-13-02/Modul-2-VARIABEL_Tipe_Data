package main

import (
	"fmt"
	// "math"
)

func LuasLingkaran(r float64) float64 {
	const pi float64 = 3.14159
	// return math.Pi * (r * r)
	return pi * (r * r)
}

func main() {
	var r float64

	fmt.Print("Silahkan Masukkan Jari Jari Lingkaran (r): ")
	fmt.Scan(&r)

	result := LuasLingkaran(r)

	fmt.Printf("Luas Lingkaran: %.1f\n", result)
}
