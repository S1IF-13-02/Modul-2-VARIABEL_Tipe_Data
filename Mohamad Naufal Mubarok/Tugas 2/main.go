package main

import (
	"fmt"
)

func main() {
	var nama, nim, kelas string

	// Input pekenalan
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim)

	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&kelas)

	// Output perkenalan
	fmt.Println("Perkenalkan saya adalah", nama, 
		", salah satu mahasiswa Prodi S1-INFORMATIKA dari kelas", kelas, 
		"dengan NIM", nim,".")
}
