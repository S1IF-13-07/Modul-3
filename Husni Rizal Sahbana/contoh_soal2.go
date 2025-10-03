package main

import "fmt"

func main() {

	var panjang_alas, tinggi_segitiga, luas_segitiga float64

	fmt.Print("Masukan panjang alas segitiga : ")
	fmt.Scan(&panjang_alas)

	fmt.Print("Masukan tinggi segitiga : ")
	fmt.Scan(&tinggi_segitiga)

	luas_segitiga = 0.5 * panjang_alas * tinggi_segitiga
	fmt.Print("Luasnya adalah : ", luas_segitiga)
}
