package main

import "fmt"

func main() {

	var r, phi float64

	phi = 3.14

	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scanln(&r)
	luas := phi * r * r
	fmt.Printf("Luas Lingkaran adalah %.2f", luas)
}
