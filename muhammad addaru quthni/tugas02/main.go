package main

import "fmt"

func main() {
	var (
		nama  string
		prodi string
		kelas string
		nim   int
	)

	fmt.Println("NAMA")
	fmt.Scan(&nama)

	fmt.Println("PRODI")
	fmt.Scan(&prodi)

	fmt.Println("KELAS")
	fmt.Scan(&kelas)

	fmt.Println("NIM")
	fmt.Scan(&nim)

	fmt.Println("Perekenalkan nama saya adalah", nama,
		"salah satu mahasiswa di prodi", prodi,
		"dari kelas", kelas, "dengan NIM", nim)
}
