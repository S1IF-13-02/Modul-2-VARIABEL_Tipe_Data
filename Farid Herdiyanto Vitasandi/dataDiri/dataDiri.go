package main

import "fmt"

func main() {
	
	var nama, kelas, prodi string
	var nim int

	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)	
	fmt.Print("Masukkan prodi: ")
	fmt.Scanln(&prodi)
	fmt.Print("Masukkan kelas: ")
	fmt.Scanln(&kelas)
	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim)

	fmt.Println("Perkenalkan nama saya adalah", nama + ", saya dari prodi teknik", prodi + ", kelas", kelas + " dengan NIM", nim)
}
