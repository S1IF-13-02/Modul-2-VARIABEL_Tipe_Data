package main

import "fmt"

func main(){
	var (
		nama, kelas, nim string
		temp string
		)
		fmt.Print("Masukan input string: ")
		fmt.Scanln(&nama)
		fmt.Print("Masukan input string: ")
		fmt.Scanln(&kelas)
		fmt.Print("Masukan input string: ")
		fmt.Scanln(nim)
		fmt.Println("Output awal = " + nama + " " + kelas + " " +nim)
		temp = nama
		nama = kelas
		kelas = nim
		nim = temp
		fmt.Println("Output akhir = " + nama + " " + kelas + " " +nim)
}