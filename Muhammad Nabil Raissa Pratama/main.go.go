//package main

//import "fmt"

//func main() {
//fmt.Print("hello nabil ")
//fmt.Printf("praktikum %v", 2022)
//}

//package main

//import (
//	"fmt"
//)

//func main() {
//	var a, b, c, d, e int

//	Input
//	fmt.Print("Masukkan 5 bilangan bulat: ")
//	fmt.Scan(&a, &b, &c, &d, &e)

//	Proses
//	hasil := a + b + c + d + e

//	fmt.Println("Hasil penjumlahan =", hasil)
//}

//package main

//import "fmt"

//func main() {
//	var (
//		satu, dua, tiga string
//		temp            string
//	)
//	fmt.Print("masukan input string :")
//	fmt.Scanln(&satu)
//	fmt.Print("masukan input string :")
//	fmt.Scanln(&dua)
//	fmt.Print("masukan input string :")
//	fmt.Scanln(&tiga)
//	fmt.Println("output awal =" + satu + "" + dua + "" + tiga)
//	temp = satu
///	satu = dua
//	dua = tiga
//	tiga = temp
//	fmt.Println("output akhir =" + satu + "" + dua + "" + tiga)
//}

//package main

//import "fmt"

//func main() {
//	var (
//		nama, nim, kelas string
//	)
//	fmt.Print("Masukan Nama :")
//	fmt.Scanln(&nama)
//	fmt.Print("Masukan Nim :")
//	fmt.Scanln(&nim)
//	fmt.Print("Masukan Kelas:")
//	fmt.Scanln(&kelas)

//	fmt.Println("Perkenalkan Nama Saya "+nama+"", "salah satu mahasiswa prodi S1 IF dari "+kelas+" dengan NIM "+""+nim)

//}

//package main

//import "fmt"

//func main() {
//	var r float64
//	fmt.Print("masukan jari jari lingkaran")
//	fmt.Scan(&r)

//	luas := 3.14 * r * r
//	fmt.Printf("luas lingkaran adalah %.2f\n", luas)

//}

//package main

//import "fmt"

//func main() {
//	var C, F float64

//	fmt.Print("input suhu dalam farenheit")
//	fmt.Scanln(&F)
//	C = (F - 32) * 5 / 9
//	fmt.Printf("berarti suhui didalam celcius  : %.2f", C)

//}

//package main

//import "fmt"

//func main() {
//	fmt.Println("hello worl")

//}

//package main

//import "fmt"

//func main() {
//	fmt.Println("hai")
//}

//package main

//import "fmt"

//func main() {
//	var x int = 10
//	var y float64 = float64(x) // casting dari int ke float64

//	fmt.Println("x =", x)
//	fmt.Println("y =", y)
//}

//package main

//import (
//	"fmt"
//	"strconv"
//)

//func main() {
//	var num int = 123
//	var str string = strconv.Itoa(num) // int → string

//	fmt.Println("num =", num)
//	fmt.Println("str =", str)
//}

//package main

//import (
//	"fmt"
//	"strconv"
//)

//func main() {
//	str := "123"                  // string angka
//	num, err := strconv.Atoi(str) // ubah ke int

//	if err != nil {
//		fmt.Println("Error:", err)
//	} else {
//		fmt.Println("Hasil int:", num)
//	}

//package main

//import "fmt"

//func main() {
//	var a int = 10
//	var b int = 3

//	hasil := a % b // 10 mod 3

//		fmt.Println("Sisa dari", a, "bagi", b, "=", hasil)
//	}
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var tahun string = "170845"
	tahunString, _ := strconv.Atoi(tahun)
	fmt.Println(tahunString + 1000)

	var bilangan int = tahunString + 15
	var teksbilangan string = strconv.Itoa(bilangan)
	fmt.Println(teksbilangan)
}
