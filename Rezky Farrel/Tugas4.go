package main

import "fmt"

func Tugas4() {
	var (
		F, C float64
	)
	fmt.Print("Masukan suhu dalam fahrenheit: ")
	fmt.Scanln(&F)
	C = (F - 32) * 5 / 9
	fmt.Printf("Suhu dalam celcius adalah %.2f\n", C)
}
