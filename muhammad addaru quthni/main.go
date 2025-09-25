package main

import "fmt"

func main() {
	var strs [3]string
	fmt.Print("input string: ")
	fmt.Scanln(&strs[0])
	fmt.Print("input string: ")
	fmt.Scanln(&strs[1])
	fmt.Print("input string: ")
	fmt.Scanln(&strs[2])
	fmt.Println("Output awal =", strs[0], strs[1], strs[2])
	strs[0], strs[1], strs[2] = strs[1], strs[2], strs[0]
	fmt.Println("Output akhir =", strs[0], strs[1], strs[2])
}
