package main

import (
	"fmt"
	"math"
)

func main() {
	// Data input jari-jari
	radiusData := []float64{10, 15, 25}

	fmt.Println("No\tJari-jari\tLuas")
	for i, r := range radiusData {
		luas := math.Pi * r * r
		fmt.Printf("%d\t%.1f\t\t%.1f\n", i+1, r, luas)
	}
}
