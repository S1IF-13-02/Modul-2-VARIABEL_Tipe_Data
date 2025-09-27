package main

import (
	"fmt"
)

func main() {
	var jariJari float64
	const pi = 3.14

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&jariJari)

	luas := pi * jariJari * jariJari

	fmt.Printf("Luas lingkaran = %.1f\n", luas)
}
