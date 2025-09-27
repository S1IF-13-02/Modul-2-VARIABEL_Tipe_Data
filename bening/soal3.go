package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&r)
	luas := math.Pi * math.Pow(r, 2)

	fmt.Printf("Luas lingkaran adalah: %.1f\n", luas)

}
