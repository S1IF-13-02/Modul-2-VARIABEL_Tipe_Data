package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan NIM: ")
	nim, _ := reader.ReadString('\n')
	nim = strings.TrimSpace(nim)

	fmt.Print("Masukkan kelas: ")
	kelas, _ := reader.ReadString('\n')
	kelas = strings.TrimSpace(kelas)

	fmt.Println("\nOutput:")
	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s.\n", nama, kelas, nim)
}
