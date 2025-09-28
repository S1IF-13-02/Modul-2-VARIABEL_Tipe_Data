package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari float64

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&jariJari)

	luas := math.Pi * jariJari * jariJari

	fmt.Printf("Luas lingkaran adalah: %.1f\n", luas)
}
