package main

import "fmt"

func main() {

	var (
		F float64
		C float64
	)

	fmt.Print("Masukan suhu (F) : ")
	fmt.Scanln(&F)

	C = (F - 32) * 5 / 9

	fmt.Printf("Suhu dalam (C): %.0f", C)
}
