package main

import "fmt"

func main() {
	var nama, nim, kelas string
	fmt.Print("masukan nama:")
	fmt.Scan(&nama)
	fmt.Print("masukan nim:")
	fmt.Scan(&nim)
	fmt.Print("masukan kelas:")
	fmt.Scan(&kelas)
	fmt.Println("perkenalkan nama saya adalah", nama, "salah satu mahasiswa prodi SI-IF dari kelas", kelas, "dengan NIM", nim, ".")
}
