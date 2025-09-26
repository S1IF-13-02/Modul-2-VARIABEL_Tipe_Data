package main

import (
	"fmt"
	"math"
)

func main() {
	// daftar jari-jari
	radii := []float64{7, 14, 20}

	// cetak header
	fmt.Printf("%-5s %-10s %-10s\n", "No", "Masukan", "Keluaran")

	// looping untuk hitung luas tiap jari-jari
	for i, r := range radii {
		luas := math.Pi * r * r
		fmt.Printf("%-5d %-10.0f %-10.1f\n", i+1, r, luas)
	}
}
