package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	// Input jari-jari
	fmt.Scanln(&r)

	// Hitung luas lingkaran
	luas := math.Pi * r * r

	// Tampilkan hasil dengan 1 angka di belakang koma
	fmt.Printf("%.1f\n", luas)
}