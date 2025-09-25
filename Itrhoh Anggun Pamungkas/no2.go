package main

import "fmt"

func main() {
	var nama, nim, kelas string

	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)

	fmt.Print("Masukkan Kelas: ")
	fmt.Scan(&kelas)

	fmt.Println("Perkenalkan saya adalah", nama, ", salah satu mahasiswa Prodi S1-IF dari kelas", kelas, "dengan NIM", nim, ".")
}