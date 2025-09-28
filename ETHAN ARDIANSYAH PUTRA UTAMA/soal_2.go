package main

import "fmt"

func main() {
	var nama, nim, kelas string

	fmt.Scan(&nama, &nim, &kelas)

	fmt.Printf("\n Perkenalkan saya adalah %s mahasiswa dari program studi S1-IF dengan nim %s dari kelas %s. \n", nama, nim, kelas)
}
