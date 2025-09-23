package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Membuat reader untuk membaca input dari pengguna
	reader := bufio.NewReader(os.Stdin)

	// Input Nama
	fmt.Print("Masukkan Nama: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	// Input NIM
	fmt.Print("Masukkan NIM: ")
	nim, _ := reader.ReadString('\n')
	nim = strings.TrimSpace(nim)

	// Input Kelas
	fmt.Print("Masukkan Kelas: ")
	kelas, _ := reader.ReadString('\n')
	kelas = strings.TrimSpace(kelas)

	// Menampilkan Resume
	fmt.Println("\n=== Resume Mahasiswa ===")
	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s.\n", nama, kelas, nim)
}
