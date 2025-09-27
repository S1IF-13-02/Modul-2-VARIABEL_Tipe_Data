package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	fmt.Scan(&radius)

	luas := math.Pi * radius * radius

	fmt.Printf("jari-jari= %.1f maka luas= %.1f\n", radius, luas)
}
