package main

import "fmt"

func main() {
	var F, C float64

	fmt.Print("masukan suhu dalam F: ")
	fmt.Scan(&F)

	C = (F - 32) * 5 / 9

	fmt.Printf("suhu %.0F", C)
}
