package main

import "fmt"

func main() {

	var (
		π, j float64
	)

	π = 3.14

	fmt.Print("Masukan jari jari :")
	fmt.Scanln(&j)

	luas := π * j * j

	fmt.Printf("Luas lingkaran : %.2f\n", luas)

}
