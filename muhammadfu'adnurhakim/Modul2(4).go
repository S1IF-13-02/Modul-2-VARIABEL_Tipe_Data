package main

import "fmt"

func main() {
	var fahrenheit float64

	fmt.Print("Diketahui suhu Fahrenheit : ")
	fmt.Scanln(&fahrenheit)

	celcius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("Maka %.f°F : %.f°C\n", fahrenheit, celcius)

}
