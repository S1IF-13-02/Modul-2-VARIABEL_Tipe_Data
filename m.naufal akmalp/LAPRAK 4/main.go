package main

import "fmt"

func main() {
	var (
		F, C float64
	) 
	C=(F-32) * 5/9
	fmt.Print("Masukan suhu dalam fahrenhait: ")
	fmt.Scanln(&F)
	fmt.Printf("Suhu dalam celcius adalah %2f\n", C)
	
}