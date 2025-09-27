package main

import (
	"fmt"
)

func main() {
	var fahrenheit int
	var celsius float64

	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	celsius = float64(fahrenheit-32) * 5.0 / 9.0

	fmt.Printf("Suhu dalam Celcius = %.0f\n", celsius)
}
