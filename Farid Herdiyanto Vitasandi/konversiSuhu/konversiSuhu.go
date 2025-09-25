package main

import "fmt"

func main() {

	var f int

	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scan(&f)

	C := (f - 32) * 5 / 9

	fmt.Println("Suhu dalam Celcius adalah: ",C)
}
