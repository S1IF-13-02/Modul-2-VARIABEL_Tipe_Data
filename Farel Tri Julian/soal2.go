package main

import "fmt"

func InputBiodata(nama string, nim string, kelas string) string {
	return "Perkenalkan nama saya adalah" + " " + nama + "," + " " + "salah satu\nmahasiswa Prodi S1-IF dari kelas" + " " + kelas + "\ndengan NIM" + " " + nim
}

func main() {
	var nama, nim, kelas string

	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim)

	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&kelas)

	result := InputBiodata(nama, nim, kelas)

	fmt.Println(result)
}
