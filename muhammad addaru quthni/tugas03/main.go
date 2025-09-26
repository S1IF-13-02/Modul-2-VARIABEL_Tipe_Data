package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&radius)

	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", radius, area)
}
