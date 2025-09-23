package main

import "fmt"

func main () {
	var r, l float64

	fmt.Print("Masukan Jari-Jari: ")
	fmt.Scan(&r)

	pi := 3.14

	l = pi * r * r 

	fmt.Printf("Luas lingkaran: %.1f\n", l)
}