package main

import (
	"fmt"
)

func main() {
	// daftar suhu Fahrenheit
	fahrenheit := []float64{32, 77, 212}

	// cetak header
	fmt.Printf("%-5s %-10s %-10s\n", "No", "Masukan", "Keluaran")

	// looping untuk konversi ke Celcius
	for i, f := range fahrenheit {
		c := (f - 32) * 5 / 9
		fmt.Printf("%-5d %-10.0f %-10.0f\n", i+1, f, c)
	}
}
