package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64 // variabel untuk jari-jari

	// Minta input jari-jari dari user
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&r)

	// Rumus luas lingkaran = π * r * r
	luas := math.Pi * r * r

	// Tampilkan hasil
	fmt.Printf("Luas lingkaran = %.1f\n", luas)
}
