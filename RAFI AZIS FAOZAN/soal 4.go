package main

import "fmt"

func main() {
	var fahrenheit int
	var celsius float64

	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	celsius = float64(fahrenheit-32) * 5 / 9

	fmt.Printf("Suhu dalam Celsius: %.2f°C\n", celsius)
}
