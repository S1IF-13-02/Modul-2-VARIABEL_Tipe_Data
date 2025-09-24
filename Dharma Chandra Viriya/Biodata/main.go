package main

import "fmt"

func ShowBiodata(nama, nim, prodi, kelas string) {
	fmt.Printf("Perkenalkan nama saya adalah %s salah satu\nmahasiswa Prodi %s dari kelas %s\ndengan NIM %s\n", nama, nim, prodi, kelas)
}

func main() {
	var nama, nim, prodi, kelas string

	fmt.Print("Silahkan Masukkan Nama Anda: ")
	fmt.Scanln(&nama)

	fmt.Print("Silahkan Masukkan NIM Anda: ")
	fmt.Scanln(&nim)

	fmt.Print("Silahkan Masukkan Prodi Anda: ")
	fmt.Scanln(&prodi)

	fmt.Print("Silahkan Masukkan Kelas Anda: ")
	fmt.Scanln(&kelas)

	ShowBiodata(nama, nim, prodi, kelas)
}
