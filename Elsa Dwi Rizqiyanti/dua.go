package main

import (
	"fmt"
)

func main() {
	// Data langsung diisi
	nama := "Elsa Dwi Rizqiyanti"
	nim := "109082500090"
	kelas := "S1-IF-02"

	// Output biodata
	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s.\n",
		nama, kelas, nim)
}
