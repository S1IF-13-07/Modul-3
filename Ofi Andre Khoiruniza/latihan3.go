package main

import "fmt"

func main() {
	var rupiah, dolar int
	fmt.Print("Masukan: ")
	fmt.Scan(&rupiah)
	dolar = rupiah / 15000
	fmt.Println("Hasilnya =", dolar)
}
