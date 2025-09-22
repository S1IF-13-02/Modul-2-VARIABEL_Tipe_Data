package main

import (
	"fmt"
)

func main() {
	var (
		naufal, ganteng, banget string
		temp                    string
	)

	// Input
	fmt.Print("Masukkan input pertama: ")
	fmt.Scanln(&naufal)
	fmt.Print("Masukkan input kedua: ")
	fmt.Scanln(&ganteng)
	fmt.Print("Masukkan input ketiga: ")
	fmt.Scanln(&banget)

	// Tampilkan output awal
	fmt.Println("Output awal = " + naufal + " " + ganteng + " " + banget)

	// Proses pertukaran (rotasi ke kiri)
	temp = naufal
	naufal = ganteng
	ganteng = banget
	banget = temp

	// Tampilkan output akhir
	fmt.Println("Output akhir = " + naufal + " " + ganteng + " " + banget)
}
