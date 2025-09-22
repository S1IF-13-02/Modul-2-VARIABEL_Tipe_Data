package main

import (
	"fmt"
)

func main() {
	var F float64 // Fahrenheit

	// Input suhu Fahrenheit
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&F)

	// Celcius
	C := (F - 32) * 5 / 9

	// Output hasil
	fmt.Printf("Suhu dalam Celcius = %.0f\n", C)
}
