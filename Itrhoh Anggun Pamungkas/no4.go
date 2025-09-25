package main
import "fmt"

func main() {
	var F, C float64
	
	fmt.Print("Masukkan suhu dalam F: ")
	fmt.Scan(&F)
	
	C = (F - 32) * 5 / 9
	
	fmt.Printf("Suhu %.0f°F sama dengan %.0f°C\n", F, C)
}