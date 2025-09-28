package main

import "fmt"

func main() {
	var nama string
	var nim string
	var kelas string

	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim)
	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&kelas)

	fmt.Println("Perkenalkan nama saya adalah " + nama + ", salah satu mahasiswa Prodi S1-IF dari kelas " + kelas + " dengan NIM " + nim + ("."))

}
