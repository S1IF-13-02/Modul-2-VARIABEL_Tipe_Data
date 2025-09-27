package main

import (
	"fmt"
)

func main() {
	var (
		farenheit float64
		celcius float64
	)

	fmt.Print("Masukkan suhu dalam Farenheit =  ")
	fmt.Scanln(&farenheit)

	// farenheit =  c * 9.0 / 5.0 + 32	
	celcius = (farenheit - 32) * 5.0 / 9.0

	// celcius = farenheit - 32 
	
	fmt.Println("Suhu dalam Celcius = ", celcius)
}