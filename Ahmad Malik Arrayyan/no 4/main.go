package main

import "fmt"

func KonversiSuhuFahrenheitToCelsius(f int) int {
	return ((f - 32) * 5) / 9
}

func main(){
	var f int

		fmt.Print("Masukkan f: ")
		fmt.Scan(&f)

		result := KonversiSuhuFahrenheitToCelsius(f)

		fmt.Printf("Celsius: %d\n", result)
}