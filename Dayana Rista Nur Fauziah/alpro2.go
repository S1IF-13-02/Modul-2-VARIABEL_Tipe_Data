package main

import "fmt"

func main() {
	var nama, nim, kelas, kampus string

	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim)

	fmt.Print("Masukkan kelas: ")
	fmt.Scanln(&kelas)

	fmt.Print("Masukkan asal kampus: ")
	fmt.Scanln(&kampus)

	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa dari %s, kelas %s dengan NIM %s.\n", nama, kampus, kelas, nim)
}
