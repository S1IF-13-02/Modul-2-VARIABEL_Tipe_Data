package main

import "fmt"

func main() {
	var nama, nim, kelas string

	fmt.Scan(&nama, &nim, &kelas)

	fmt.Printf("perkenalkan saya adalah %s, salah satu mahasiswa prodi S1 - IF dari kelas %s dengan nim %s.", nama, nim, kelas)
}