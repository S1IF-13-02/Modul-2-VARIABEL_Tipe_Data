package main

import "fmt"

func main() {
	var Nama string
	var Kelas string
	var NIM string

	fmt.Println("Biodata Mahasiswa")
	fmt.Print("Masukkan Nama:")
	fmt.Scanln(&Nama)
	fmt.Print("Masukkan Kelas:")
	fmt.Scanln(&Kelas)
	fmt.Print("Masukkan NIM:")
	fmt.Scanln(&NIM)

	fmt.Printf("Perkenalkan nama saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s.\n",
		Nama, Kelas, NIM)

}
