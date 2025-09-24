package main

import (
	"fmt"
)

func main() {
	// Data input Fahrenheit
	fahrenheitData := []int{50, 100, 150}

	fmt.Println("No\tFahrenheit\tCelcius")
	for i, f := range fahrenheitData {
		c := (f - 32) * 5 / 9
		fmt.Printf("%d\t%d\t\t%d\n", i+1, f, c)
	}
}
