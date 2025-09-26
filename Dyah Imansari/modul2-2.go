package main

import (
	"fmt"
)

func main() {
	var (
		nama, nim, kelas string
	)
	fmt.Print("Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Nim: ")
	fmt.Scanln(&nim)
	fmt.Print("Kelas: ")
	fmt.Scanln(&kelas)

	fmt.Println("Perkenalkan saya adalah " + nama + ", salah satu mahasiswa Prodi S1-IF dari kelas " + kelas + " dengan nim " + nim)

}
