package main

import "fmt"
func main() {
	var F, C float64
	fmt.Print("Masukan suhu dalam Farenheit: ")
	fmt.Scanln(&F)
	C = (F - 32) * 5 / 9
	fmt.Printf("Berarti Suhu di Celcius : %.2f", C)
}