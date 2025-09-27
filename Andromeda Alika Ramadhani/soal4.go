package main

import "fmt"

func main() {
	var suhu_fahreinheit int

	fmt.Scan(&suhu_fahreinheit)

	suhu_celcius := 9 / 5 * (32 - suhu_fahreinheit)

	fmt.Printf("suhu_celcius: %d\n", suhu_celcius)
}
