package main

import "fmt"

func main() {
	var F, C float64
	fmt.Print("F = ")
	fmt.Scan(&F)
	C = (F - 32) * 5 / 9
	fmt.Println("nilai C =", C, "°C")

}
