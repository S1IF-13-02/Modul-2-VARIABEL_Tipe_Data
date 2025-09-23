package main

import "fmt"

func Tugas3() {
	var (
		p, r float64
	)
	p = 3.14
	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scanln(&r)
	luas := p * r * r
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", r, luas)
}
