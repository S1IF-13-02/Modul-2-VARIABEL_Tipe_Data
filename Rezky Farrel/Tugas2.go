package main

import "fmt"

func Tugas2() {
	var (
		nama, nim, kelas string
	)
	fmt.Print("Masukan nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukan nim: ")
	fmt.Scanln(&nim)
	fmt.Print("Masukan kelas: ")
	fmt.Scanln(&kelas)

	fmt.Println("Nama saya adalah " + nama + ", nim saya " + nim + ", kelas saya " + kelas)

}

func main() {
	Tugas2()
}
