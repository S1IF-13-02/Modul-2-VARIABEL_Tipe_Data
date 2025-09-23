package main

import "fmt"

func main() {
	var (
		apel, anggur, tempe string
		temp            string
	)
	fmt.Print("Masukan input string1: ")
	fmt.Scanln(&apel)
	fmt.Print("Masukan input string2: ")
	fmt.Scanln(&anggur)
	fmt.Print("Masukan input string3: ")
	fmt.Scanln(&tempe)
	fmt.Println("Output awal = " + apel + " " + anggur + " " + tempe)
	temp = apel
	apel = anggur
	anggur = tempe
	tempe = temp
	fmt.Println("Output akhir = " + apel + " " + anggur + " " + tempe)
}
