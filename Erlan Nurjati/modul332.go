package main

import "fmt"

func main() {
	var alas, tinggi, luas float64
	fmt.Print("Masukkan alas: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)
	luas = 0.5 * alas * tinggi
	fmt.Print("Jadi luas segitiga anda adalah: ", luas)
}
