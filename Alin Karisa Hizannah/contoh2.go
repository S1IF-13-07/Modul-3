package main

import "fmt"

func main() {
	var alas, tinggi, luas float64
	fmt.Print("Masukkan Alas: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan Tinggi: ")
	fmt.Scan(&tinggi)
	luas = 0.5 * alas * tinggi

	fmt.Printf("Luas Segitiga = %.f\n", luas)
}
