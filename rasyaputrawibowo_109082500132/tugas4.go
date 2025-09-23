package main

import "fmt"

func tugas4() {
	var (
		F, C float64
	) 
	C=(F-32) * 5/9
	fmt.Print("Masukan suhu dalam fahrenhait: ")
	fmt.Scanln(&F)
	fmt.Printf("Suhu dalam celcius adalah %2f\n", C)
}

func main() {
	tugas4()
}