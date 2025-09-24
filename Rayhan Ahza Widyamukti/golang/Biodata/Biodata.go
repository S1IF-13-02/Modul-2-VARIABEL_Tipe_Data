package main

import "fmt"

func main() {
	var nama, prodi, kelas, nim string

	fmt.Print("Masukkan nama : ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan prodi : ")
	fmt.Scanln(&prodi)
	fmt.Print("Masukkan kelas : ")
	fmt.Scanln(&kelas)
	fmt.Print("Masukkan NIM : ")
	fmt.Scanln(&nim)

	fmt.Println("\nPerkenalkan saya adalah", nama+", salah satu")
	fmt.Println("mahasiswa Prodi", prodi, "dari kelas", kelas)
	fmt.Println("dengan NIM", nim+".")
}
