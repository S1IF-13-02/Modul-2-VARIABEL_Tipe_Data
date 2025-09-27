package main

import (
	"fmt"
	"math"
)


func main() {
	var (
		r, total float64
	)

	// pi = 22.0 / 7.0

	fmt.Print("Masukkan jari jari lingkaran =  ")
	fmt.Scanln(&r)

	total =  math.Pi * r * r
	
	fmt.Println("Luas lingkaran adalah", total)
}