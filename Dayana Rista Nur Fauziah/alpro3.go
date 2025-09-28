package main

import (
	"fmt"
)

func main() {
	var fahrenheit float64 = 77

	celcius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("Suhu dalam Fahrenheit: %.0f°F\n", fahrenheit)
	fmt.Printf("Suhu dalam Celcius: %.0f°C\n", celcius)
}
