package main

import "fmt"

func main() {
	var r float64
	const PI = 3.14159

	fmt.Print("masukan jari-jari lingkaran: ")
	fmt.Scan(&r)

	luas := PI * r * r

	fmt.Printf("luas lingkaran adalah %.1f\n", luas)
}
