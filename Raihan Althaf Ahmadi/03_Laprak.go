package main

import "fmt"
func main() {
	var r, π float64
	π = 3.14
	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scanln(&r)
	luas :=  π * r * r 
	fmt.Printf("Luas Lingkaran adalah %.2f", luas)
}