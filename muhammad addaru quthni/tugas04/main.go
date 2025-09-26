package main

import (
    "fmt"
)

func fahrenheitToCelsius(f int) float64 {
    return float64(f-32) * 5.0 / 9.0
}

func main() {
    var f int
    fmt.Print("Masukkan suhu dalam Fahrenheit: ")
    fmt.Scanln(&f)

    c := fahrenheitToCelsius(f)
    fmt.Printf("Suhu dalam Celcius: %.2f\n", c)
}