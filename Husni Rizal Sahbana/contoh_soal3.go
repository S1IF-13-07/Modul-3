package main

import "fmt"

func main() {
	var rupiah int
	var dolar = 15000
	fmt.Print("Masukan nominal uang rupiah : ")
	fmt.Scan(&rupiah)
	hasil := rupiah / dolar
	fmt.Print("Hasilnya adalah : ", hasil)
}
