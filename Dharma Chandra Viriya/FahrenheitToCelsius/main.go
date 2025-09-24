package main

import "fmt"

func FahrenheitToCelsius(f int) int {
	return ((f - 32) * 5) / 9
}

func main() {
	var f int

	fmt.Print("Silahkan Masukkan Suhu Fahrenheit (f): ")
	fmt.Scan(&f)

	result := FahrenheitToCelsius(f)

	fmt.Printf("%d Fahrenheit -> %d Celsius\n", f, result)
}
