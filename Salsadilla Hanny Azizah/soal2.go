package main

import "fmt"

func main() {
	var (
		nama, kelas string
		nim int
	)
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan kelas: ")
	fmt.Scanln(&kelas)
	
	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim)
	
	fmt.Println("Perkenalkan saya adalah", nama +", salah satu mahasiswa dari Prodi", kelas+ ", dengan NIM", nim )

}