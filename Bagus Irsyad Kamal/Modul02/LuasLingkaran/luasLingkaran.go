package main

import "fmt"

func main() {
	var r float64
	const pi = 3.14159

	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scan(&r)

	luas := pi * r * r

	fmt.Printf("Luas lingkaran = %.1f\n", luas)

}
