package main

import "fmt" 

func main() {
	var alas float64
	var tinggi float64
	var luas float64
	fmt.Print("Masukkan alas segitiga: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)
	luas = 0.5 * alas * tinggi
	fmt.Println("Luas segitiga adalah:", luas)
}