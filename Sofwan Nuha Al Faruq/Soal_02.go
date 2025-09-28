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

 	fmt.Printf("\nPerkenalkan saya adalah %s, salah satu mahasiswa Prodi %s dari kelas %s dengan NIM %s.\n",
        nama, prodi, kelas, nim)

}
