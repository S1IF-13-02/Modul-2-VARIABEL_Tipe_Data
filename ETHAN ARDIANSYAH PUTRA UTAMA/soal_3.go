package main

import "fmt"

func main() {

	const phi = 3.14
	var r float64

	fmt.Print("Masukkan panjang jari-jari: ")
	fmt.Scan(&r)

	luas := phi * r * r

	fmt.Println("Luas lingkaran adalah: ", luas)

}
