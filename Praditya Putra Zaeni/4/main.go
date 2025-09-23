package main

import (
	"fmt"
)

func main() {
	var f float64

	// Input suhu dalam Fahrenheit
	fmt.Scanln(&f)

	// Konversi ke Celcius
	c := (f - 32) * 5 / 9

	// Output Celcius (dibulatkan tanpa angka desimal)
	fmt.Printf("%.0f\n", c)
}
