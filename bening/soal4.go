package main

import "fmt"

func main() {
	var Fahrenheit float64

	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&Fahrenheit)
	celcius := (Fahrenheit - 32) * 5.0 / 9.0
	fmt.Printf("\nSuhu Fahrenheit(F) = %.0f\n", Fahrenheit)
	fmt.Printf("Suhu Celcius (C) = %.2f\n", celcius)
}
